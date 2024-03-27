package gosdk

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	api "github.com/begonia-org/go-sdk/api/v1"
	common "github.com/begonia-org/go-sdk/common/api/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

type BegoniaClient interface {
}

type BegoniaClientImpl struct {
	cli     *http.Client
	baseUrl string
	signer  AppAuthSigner
}

const UPLOAD_API = "/api/v1/file"
const INIT_PART_API = "/api/v1/file/part/init"
const UPLOAD_PART_API = "/api/v1/file/part"
const COMPLETE_PART_API = "/api/v1/file/part/complete"
const ABORT_PART_API = "/api/v1/file/part/abort"
const Download_API = "/api/v1/file"
const Metadata_API = "/api/v1/file/metadata"
const Download_PART_API = "/api/v1/file/part"

func NewBegoniaClient(addr, accessKey, secretKey string) *BegoniaClientImpl {
	return &BegoniaClientImpl{
		cli:     &http.Client{},
		baseUrl: addr,
		signer:  NewAppAuthSigner(accessKey, secretKey),
	}
}
func (bc *BegoniaClientImpl) requestSignature(_ context.Context, req *http.Request) error {
	// TODO
	header := NewRequestHeader()
	for k, v := range req.Header {

		req.Header.Set(k, strings.Join(v, ","))
	}
	gw := &GatewayRequest{
		Headers: header,
		Method:  req.Method,
		URL:     req.URL,
		Host:    req.Host,
		Payload: req.Body,
	}
	err := bc.signer.SignRequest(gw)
	if err != nil {
		return err
	}
	for _, k := range gw.Headers.Keys() {
		v := gw.Headers.Get(k)
		if strings.Contains(v, ",") {
			values := strings.Split(v, ",")
			for _, value := range values {
				req.Header.Add(k, value)
			}
		} else {
			req.Header.Set(k, v)

		}

	}
	return nil

}
func (bc *BegoniaClientImpl) Get(ctx context.Context, uri string, headers map[string]string) (*http.Response, error) {
	api, _ := url.JoinPath(bc.baseUrl, uri)
	req, _ := http.NewRequest(http.MethodGet, api, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	err := bc.requestSignature(ctx, req)
	if err != nil {
		return nil, err
	}
	return bc.cli.Do(req)
}

func (bc *BegoniaClientImpl) Post(ctx context.Context, uri string, headers map[string]string, payload io.Reader) (*http.Response, error) {
	api, _ := url.JoinPath(bc.baseUrl, uri)
	req, _ := http.NewRequest(http.MethodPost, api, payload)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	err := bc.requestSignature(ctx, req)
	if err != nil {
		return nil, err
	}
	return bc.cli.Do(req)
}

func (bc *BegoniaClientImpl) UploadFile(ctx context.Context, filePath string) (string, error) {
	api, _ := url.JoinPath(bc.baseUrl, UPLOAD_API)
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	req, _ := http.NewRequest(http.MethodPost, api, file)
	// for k, v := range headers {
	// 	req.Header.Set(k, v)
	// }
	err = bc.requestSignature(ctx, req)
	if err != nil {
		return "", err
	}
	rsp, err := bc.cli.Do(req)
	if err != nil {
		return "", err
	}
	if rsp.StatusCode > 400 {
		return "", errors.New("upload file failed")
	}
	if rsp != nil {
		defer rsp.Body.Close()
		apiRsp := &common.HttpResponse{}
		data, _ := io.ReadAll(rsp.Body)
		err := protojson.Unmarshal(data, apiRsp)
		if err != nil {
			return "", err
		}
		val := apiRsp.Data.AsMap()
		if status := apiRsp.Code; common.Code(status) != common.Code_OK {
			return "", errors.New(apiRsp.Message)

		}
		return val["uri"].(string), nil
	}
	return "", nil
}

func (bc *BegoniaClientImpl) CreateEndpoint(ctx context.Context, endpoint *api.AddEndpointRequest) (*api.AddEndpointResponse, error) {
	apiEndpoint, _ := url.JoinPath(bc.baseUrl, "/api/v1/endpoint")
	payload, err := protojson.Marshal(endpoint)
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest(http.MethodPost, apiEndpoint, strings.NewReader(string(payload)))
	req.Header.Set("Content-Type", "application/json")
	err = bc.requestSignature(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp, err := bc.cli.Do(req)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode > 400 {
		return nil, errors.New("create endpoint failed")
	}
	if rsp != nil {
		defer rsp.Body.Close()
		apiRsp := &common.HttpResponse{}
		data, _ := io.ReadAll(rsp.Body)
		err := protojson.Unmarshal(data, apiRsp)
		if err != nil {
			return nil, err
		}
		val := apiRsp.Data
		if status := apiRsp.Code; common.Code(status) != common.Code_OK {
			return nil, errors.New(apiRsp.Message)
		}
		jsonBytes, err := protojson.Marshal(val)
		if err != nil {
			return nil, fmt.Errorf("Failed to marshal data to JSON: %w", err)
		}

		// 将 JSON 字符串反序列化为目标 pb 结构体
		var endpoint *api.AddEndpointResponse
		if err := protojson.Unmarshal(jsonBytes, endpoint); err != nil {
			return nil, fmt.Errorf("Failed to unmarshal JSON to Endpoint: %w", err)
		}
		return endpoint, nil
	}
	return nil, nil
}
