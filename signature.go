// BEGONIA API Gateway Signature
// based on https://github.com/datastream/aws/blob/master/signv4.go
package gosdk

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"google.golang.org/grpc/metadata"
)

type RequestHeader struct {
	headers map[string]string
}

type GatewayRequest struct {
	Headers *RequestHeader
	Payload io.ReadCloser
	Method  string
	URL     *url.URL
	Host    string
}
type AppAuthSigner interface {
	Sign(request *GatewayRequest) (string, error)
	SignRequest(request *GatewayRequest) error
}

type AppAuthSignerImpl struct {
	Key    string
	Secret string
}

const (
	DateFormat           = "20060102T150405Z"
	SignAlgorithm        = "X-Sign-HMAC-SHA256"
	HeaderXDateTime      = "X-Date"
	HeaderXHost          = "host"
	HeaderXAuthorization = "Authorization"
	HeaderXContentSha256 = "X-Content-Sha256"
	HeaderXAccessKey     = "X-Access-Key"
)

func (h *RequestHeader) Set(key, value string) {
	h.headers[strings.ToLower(key)] = value

}
func (h *RequestHeader) Get(key string) string {
	return h.headers[strings.ToLower(key)]

}
func (h *RequestHeader) Del(key string) {
	delete(h.headers, strings.ToLower(key))
}
func (h *RequestHeader) Keys() []string {
	var keys []string
	for k := range h.headers {
		keys = append(keys, k)
	}
	return keys
}
func (h *RequestHeader) ToMetadata() metadata.MD {
	md := metadata.New(h.headers)
	return md
}
func NewRequestHeader() *RequestHeader {
	return &RequestHeader{headers: make(map[string]string)}
}

// filtersUriParams 过滤uri上的附加的参数
func filtersUriParams(xparams string, body []byte) []byte {
	if xparams == "" || body == nil {
		return body
	}
	params := strings.Split(xparams, ",")
	bodyMap := make(map[string]interface{})
	err := json.Unmarshal(body, &bodyMap)
	if err != nil {
		log.Println("json unmarshal error:", err)
		return body

	}
	for _, param := range params {
		delete(bodyMap, param)
	}
	if len(bodyMap) == 0 {
		return nil
	}
	body, err = json.Marshal(bodyMap)
	if err != nil {
		log.Println("json marshal error:", err)
		return body

	}
	return body

}
func NewGatewayRequestFromGrpc(ctx context.Context, req interface{}, fullMethod string) (*GatewayRequest, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	headers := &RequestHeader{headers: make(map[string]string)}
	host := ""
	uri := fullMethod
	method := fullMethod
	xparams := ""
	// bodySha256 := ""
	if ok {
		for k, v := range md {
			if strings.EqualFold(k, "x-forwarded-host") {
				host = v[0]
			}
			if strings.EqualFold(k, "uri") {
				uri = v[0]
			}
			if strings.EqualFold(k, "x-http-method") {
				method = v[0]
			}
			if strings.EqualFold(k, "x-gateway-params") {
				xparams = v[0]
			}

			values := []string{}
			for _, val := range v {
				hs := strings.Split(val, ",")
				for index, val := range hs {
					hs[index] = strings.TrimSpace(val)
				}
				values = append(values, hs...)
			}
			headers.Set(k, strings.Join(values, ","))
		}
	}
	u, _ := url.Parse(fmt.Sprintf("http://%s%s", host, uri))
	var payload []byte
	payload, _ = json.Marshal(req)
	if xparams != "" {
		payload = filtersUriParams(xparams, payload)
	}
	// log.Printf("payload:%s", payload)

	reader := io.NopCloser(bytes.NewBuffer(payload))

	return &GatewayRequest{Headers: headers,
		Method:  method,
		Host:    host,
		URL:     u,
		Payload: reader,
	}, nil
}
func NewGatewayRequestFromHttp(req *http.Request) (*GatewayRequest, error) {
	// headers := make(map[string]string)
	headers := &RequestHeader{headers: make(map[string]string)}
	for k, v := range req.Header {
		// 在gateway 中会被修改为application/grpc
		// 因此不参与签名
		if strings.EqualFold(k, "Content-Type") {
			continue
		}
		headers.Set(k, strings.Join(v, ","))
	}
	// payload := []byte("")
	var payload []byte = []byte("{}")
	if req.Body != nil {
		payload, _ = io.ReadAll(req.Body)
		req.Body = io.NopCloser(bytes.NewBuffer(payload))
	}
	var reader io.ReadCloser
	if payload != nil {
		reader = io.NopCloser(bytes.NewBuffer(payload))

	}
	return &GatewayRequest{Headers: headers, Method: req.Method, URL: req.URL, Host: req.Host, Payload: reader}, nil
}
func NewAppAuthSigner(key, secret string) AppAuthSigner {
	return &AppAuthSignerImpl{Key: key, Secret: secret}
}
func (app *AppAuthSignerImpl) hmacsha256(keyByte []byte, dataStr string) ([]byte, error) {
	hm := hmac.New(sha256.New, []byte(keyByte))
	if _, err := hm.Write([]byte(dataStr)); err != nil {
		return nil, err
	}
	return hm.Sum(nil), nil
}

// Build a CanonicalRequest from a regular request string
func (app *AppAuthSignerImpl) CanonicalRequest(request *GatewayRequest, signedHeaders []string) (string, error) {
	var hexencode string
	var err error
	if hex := request.Headers.Get(HeaderXContentSha256); hex != "" {
		hexencode = hex
	} else {
		bodyData, err := app.RequestPayload(request)
		if err != nil {
			return "", err
		}
		hexencode, err = app.HexEncodeSHA256Hash(bodyData)
		if err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s", strings.ToUpper(request.Method), app.CanonicalURI(request), app.CanonicalQueryString(request), app.CanonicalHeaders(request, signedHeaders), strings.Join(signedHeaders, ";"), hexencode), err
}

// CanonicalURI returns request uri
func (app *AppAuthSignerImpl) CanonicalURI(request *GatewayRequest) string {
	pattens := strings.Split(request.URL.Path, "/")
	var uriSlice []string
	for _, v := range pattens {
		uriSlice = append(uriSlice, url.PathEscape(v))
	}
	urlpath := strings.Join(uriSlice, "/")
	if len(urlpath) == 0 || urlpath[len(urlpath)-1] != '/' {
		urlpath = urlpath + "/"
	}
	// log.Printf("canonicalURI:%s", urlpath)
	return urlpath
}

// CanonicalQueryString
func (app *AppAuthSignerImpl) CanonicalQueryString(request *GatewayRequest) string {
	var keys []string
	queryMap := request.URL.Query()
	for key := range queryMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var query []string
	for _, key := range keys {
		k := escape(key)
		sort.Strings(queryMap[key])
		for _, v := range queryMap[key] {
			kv := fmt.Sprintf("%s=%s", k, escape(v))
			query = append(query, kv)
		}
	}
	queryStr := strings.Join(query, "&")
	request.URL.RawQuery = queryStr
	// // log.Printf("canonicalQueryString:%s", queryStr)
	return queryStr
}

// CanonicalHeaders
func (app *AppAuthSignerImpl) CanonicalHeaders(request *GatewayRequest, signerHeaders []string) string {
	var canonicalHeaders []string
	header := make(map[string][]string)
	for k, v := range request.Headers.headers {
		val := strings.Split(v, ",")
		header[strings.ToLower(k)] = val
	}
	for _, key := range signerHeaders {
		value := header[strings.ToLower(key)]
		if strings.EqualFold(key, HeaderXHost) {
			value = []string{request.Host}
		}

		sort.Strings(value)
		// log.Println("sort value:", value)
		canonicalHeaders = append(canonicalHeaders, key+":"+strings.TrimSpace(strings.Join(value, ",")))

	}
	return strings.Join(canonicalHeaders, "\n")
}

// SignedHeaders
func (app *AppAuthSignerImpl) SignedHeaders(r *GatewayRequest) []string {
	var signedHeaders []string
	for key := range r.Headers.headers {
		if strings.EqualFold(key, HeaderXAuthorization) {
			return app.getSignaturehHeader(r.Headers.Get(HeaderXAuthorization))
		}
		if r.Headers.headers[key] == "" {
			continue
		}
		signedHeaders = append(signedHeaders, strings.ToLower(key))
	}
	sort.Strings(signedHeaders)
	return signedHeaders
}

// RequestPayload
func (app *AppAuthSignerImpl) RequestPayload(request *GatewayRequest) (io.ReadCloser, error) {
	if request.Payload == nil {
		return io.NopCloser(bytes.NewBufferString("")), nil
	}

	var buf bytes.Buffer
	tee := io.TeeReader(request.Payload, &buf)

	// 使用 TeeReader 将数据写入缓冲区，同时返回读取器
	request.Payload = io.NopCloser(&buf)
	return io.NopCloser(tee), nil
}

// Create a "String to Sign".
func (app *AppAuthSignerImpl) StringToSign(canonicalRequest string, t time.Time) (string, error) {
	hashStruct := sha256.New()
	_, err := hashStruct.Write([]byte(canonicalRequest))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s\n%s\n%x",
		SignAlgorithm, t.UTC().Format(DateFormat), hashStruct.Sum(nil)), nil
}

// Create the  Signature.
func (app *AppAuthSignerImpl) SignStringToSign(stringToSign string, signingKey []byte) (string, error) {
	hmsha, err := app.hmacsha256(signingKey, stringToSign)
	return fmt.Sprintf("%x", hmsha), err
}

// HexEncodeSHA256Hash returns hexcode of sha256
func (app *AppAuthSignerImpl) HexEncodeSHA256Hash(body io.ReadCloser) (string, error) {
	defer body.Close()
	hashStruct := sha256.New()

	buf := make([]byte, 4096)
	for {
		n, err := body.Read(buf)
		if err != nil && err != io.EOF {
			return "", err
		}
		if n == 0 {
			break
		}
		hashStruct.Write(buf[:n])
	}

	return fmt.Sprintf("%x", hashStruct.Sum(nil)), nil
}

// Get the finalized value for the "Authorization" header. The signature parameter is the output from SignStringToSign
func (app *AppAuthSignerImpl) AuthHeaderValue(signatureStr, accessKeyStr string, signedHeaders []string) string {
	return fmt.Sprintf("%s Access=%s, SignedHeaders=%s, Signature=%s", SignAlgorithm, accessKeyStr, strings.Join(signedHeaders, ";"), signatureStr)
}

// SignRequest set Authorization header
func (app *AppAuthSignerImpl) Sign(request *GatewayRequest) (string, error) {
	unusedHeaderKeys := []string{"content-type", "content-length", "accept-encoding"}
	for _, key := range request.Headers.Keys() {
		for _, unusedKey := range unusedHeaderKeys {
			if strings.EqualFold(key, unusedKey) {
				request.Headers.Del(key)
			}
		}
	}
	var t time.Time
	var err error
	var date string
	if date = request.Headers.Get(HeaderXDateTime); date != "" {
		t, err = time.Parse(DateFormat, date)
		if err != nil {
			return "", fmt.Errorf("Failed to parse X-Date: %w", err)
		}
		if time.Since(t) > time.Minute*1 {
			return "", fmt.Errorf("X-Date is expired")
		}
	}
	if err != nil || date == "" {
		t = time.Now()
		// log.Printf("X-Date is not set, using current time,%v,%v", err, date)
		request.Headers.Set(HeaderXDateTime, t.UTC().Format(DateFormat))
	}
	request.Headers.Set(HeaderXAccessKey, app.Key)
	signedHeaders := app.SignedHeaders(request)
	// log.Printf("signedHeaders:%v", signedHeaders)

	canonicalRequest, err := app.CanonicalRequest(request, signedHeaders)
	if err != nil {
		return "", fmt.Errorf("Failed to create canonical request: %w", err)
	}
	// log.Printf("canonicalRequest:%s", canonicalRequest)
	stringToSignStr, err := app.StringToSign(canonicalRequest, t)
	// log.Printf("stringToSign:%s", stringToSignStr)
	if err != nil {
		return "", fmt.Errorf("Failed to create string to sign: %w", err)
	}
	signatureStr, err := app.SignStringToSign(stringToSignStr, []byte(app.Secret))
	// log.Printf("signature:%s", signatureStr)
	if err != nil {
		return "", fmt.Errorf("Failed to create signature: %w", err)
	}

	return signatureStr, nil
}
func (app *AppAuthSignerImpl) getSignaturehHeader(auth string) []string {
	strArr := strings.Split(auth, ",")
	for _, v := range strArr {
		if strings.Contains(strings.ToLower(v), "signedheaders") {
			signature := strings.Split(v, "=")
			return strings.Split(signature[1], ";")
		}
	}
	return nil
}
func (app *AppAuthSignerImpl) SignRequest(request *GatewayRequest) error {
	unusedHeaderKeys := []string{"content-type", "content-length", "accept-encoding"}
	for _, key := range request.Headers.Keys() {
		for _, unusedKey := range unusedHeaderKeys {
			if strings.EqualFold(key, unusedKey) {
				request.Headers.Del(key)
			}
		}
	}
	signature, err := app.Sign(request)
	if err != nil {
		return err
	}
	request.Headers.Set(HeaderXAuthorization, app.AuthHeaderValue(signature, app.Key, app.SignedHeaders(request)))
	return nil
}
