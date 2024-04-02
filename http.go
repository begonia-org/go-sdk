package gosdk

import (
	"bytes"
	"context"
	"crypto/sha256"
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

type Response struct {
	StatusCode int
	RequestId  string
	Err        error
}
type AddEndpointConfigResponse struct {
	*Response
	Id string
}
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

	gw, err := NewGatewayRequestFromHttp(req)
	if err != nil {
		return err
	}
	// log.Println("客户端开始签名")
	err = bc.signer.SignRequest(gw)
	// log.Println("客户端签名完成")
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
	req := bc.buildRequest(ctx, http.MethodGet, api, headers, nil)
	return bc.request(ctx, req)
}

func (bc *BegoniaClientImpl) Post(ctx context.Context, uri string, headers map[string]string, payload io.Reader) (*http.Response, error) {
	api, _ := url.JoinPath(bc.baseUrl, uri)
	req := bc.buildRequest(ctx, http.MethodPost, api, headers, payload)
	return bc.request(ctx, req)
}
func (bc *BegoniaClientImpl) Put(ctx context.Context, uri string, headers map[string]string, payload io.Reader) (*http.Response, error) {
	api, _ := url.JoinPath(bc.baseUrl, uri)
	req := bc.buildRequest(ctx, http.MethodPut, api, headers, payload)
	return bc.request(ctx, req)
}
func (bc *BegoniaClientImpl) Delete(ctx context.Context, uri string, headers map[string]string, payload io.Reader) (*http.Response, error) {
	api, _ := url.JoinPath(bc.baseUrl, uri)
	req := bc.buildRequest(ctx, http.MethodDelete, api, headers, payload)
	return bc.request(ctx, req)

}
func (bc *BegoniaClientImpl) Patch(ctx context.Context, uri string, headers map[string]string, payload io.Reader) (*http.Response, error) {
	api, _ := url.JoinPath(bc.baseUrl, uri)

	req := bc.buildRequest(ctx, http.MethodPatch, api, headers, payload)

	return bc.request(ctx, req)
}
func (bc *BegoniaClientImpl) buildRequest(_ context.Context, method, uri string, headers map[string]string, payload io.Reader) *http.Request {
	req, _ := http.NewRequest(method, uri, payload)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return req
}
func (bc *BegoniaClientImpl) request(ctx context.Context, req *http.Request) (*http.Response, error) {
	err := bc.requestSignature(ctx, req)
	if err != nil {
		return nil, err
	}
	return bc.cli.Do(req)
}
func (bc *BegoniaClientImpl) UploadFile(ctx context.Context, srcPath string, dst string) (string, error) {
	apiEndpoint, _ := url.JoinPath(bc.baseUrl, UPLOAD_API)

	data, err := os.ReadFile(srcPath)
	if err != nil {
		return "", err

	}
	shaer := sha256.New()
	shaer.Write(data)
	sha := shaer.Sum(nil)
	hashStr := fmt.Sprintf("%x", sha)
	content := api.UploadFileRequest{
		Key:         dst,
		Content:     data,
		Sha256:      hashStr,
		ContentType: http.DetectContentType(data),
	}
	payload, err := protojson.Marshal(&content)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal data to JSON: %w", err)

	}
	req, _ := http.NewRequest(http.MethodPut, apiEndpoint, bytes.NewReader(payload))
	// for k, v := range headers {
	// 	req.Header.Set(k, v)
	// }
	req.Header.Set("Content-Type", "application/json")
	err = bc.requestSignature(ctx, req)
	if err != nil {
		return "", fmt.Errorf("Failed to sign request: %w", err)
	}
	rsp, err := bc.cli.Do(req)
	if err != nil {
		return "", fmt.Errorf("Failed to send request: %w", err)
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

	if rsp != nil {
		defer rsp.Body.Close()
		apiRsp := &common.HttpResponse{}
		data, _ := io.ReadAll(rsp.Body)
		err := protojson.Unmarshal(data, apiRsp)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal JSON to HttpResponse: %w", err)
		}
		val := apiRsp.Data
		if status := apiRsp.Code; common.Code(status) != common.Code_OK {
			return nil, fmt.Errorf("Failed to create endpoint: %s", apiRsp.Message)
		}
		jsonBytes, err := protojson.Marshal(val)
		if err != nil {
			return nil, fmt.Errorf("Failed to marshal data to JSON: %w", err)
		}

		// 将 JSON 字符串反序列化为目标 pb 结构体
		var endpoint *api.AddEndpointResponse = &api.AddEndpointResponse{}
		if err := protojson.Unmarshal(jsonBytes, endpoint); err != nil {
			return nil, fmt.Errorf("Failed to unmarshal JSON to Endpoint: %w", err)
		}
		return endpoint, nil
	}

	return nil, errors.New("create endpoint failed")
}

func (bc *BegoniaClientImpl) RegisterEndpoint(ctx context.Context, config *api.EndpointSrvConfig) (*AddEndpointConfigResponse, error) {
	apiEndpoint, _ := url.JoinPath(bc.baseUrl, "/api/v1/endpoint/config")
	payload, err := protojson.Marshal(config)
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

	if rsp != nil {
		defer rsp.Body.Close()
		requestId := rsp.Header.Get("x-request-id")
		apiRsp := &common.HttpResponse{}
		data, _ := io.ReadAll(rsp.Body)
		err := protojson.Unmarshal(data, apiRsp)
		if err != nil {

			return &AddEndpointConfigResponse{Response: &Response{
				StatusCode: rsp.StatusCode,
				RequestId:  requestId,
				Err:        err,
			}}, err
		}
		// val := apiRsp.Data.AsMap()
		if status := apiRsp.Code; common.Code(status) != common.Code_OK {
			return &AddEndpointConfigResponse{Response: &Response{
				StatusCode: rsp.StatusCode,
				RequestId:  requestId,
				Err:        errors.New(apiRsp.Message),
			}}, errors.New(apiRsp.Message)
		}
		return nil, nil
	}

	return nil, errors.New("add endpoint config failed")

}
