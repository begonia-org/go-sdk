package client

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	gosdk "github.com/begonia-org/go-sdk"
)

type Response struct {
	StatusCode int
	RequestId  string
	Err        error
}

type BaseAPI struct {
	cli     *http.Client
	baseUrl string
	signer  gosdk.AppAuthSigner
}

const UPLOAD_API = "/api/v1/files"
const INIT_PART_API = "/api/v1/files/part/init"
const UPLOAD_PART_API = "/api/v1/files/part"
const COMPLETE_PART_API = "/api/v1/files/part/complete"
const ABORT_PART_API = "/api/v1/files/part/abort"
const Download_API = "/api/v1/files"
const Metadata_API = "/api/v1/files/metadata"
const Download_PART_API = "/api/v1/files/part"
const FILE_API = "/api/v1/files"

func NewAPIClient(addr, accessKey, secretKey string) *BaseAPI {
	return &BaseAPI{
		cli:     &http.Client{},
		baseUrl: addr,
		signer:  gosdk.NewAppAuthSigner(accessKey, secretKey),
	}
}
func (bc *BaseAPI) requestSignature(_ context.Context, req *http.Request) error {

	gw, err := gosdk.NewGatewayRequestFromHttp(req)
	if err != nil {
		return err
	}
	err = bc.signer.SignRequest(gw)
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
func (bc *BaseAPI) Get(ctx context.Context, uri string, headers map[string]string) (*http.Response, error) {
	api, _ := url.JoinPath(bc.baseUrl, uri)
	// api, _ = url.QueryUnescape(api)
	req := bc.buildRequest(ctx, http.MethodGet, api, headers, nil)
	return bc.request(ctx, req)
}

func (bc *BaseAPI) Post(ctx context.Context, uri string, headers map[string]string, payload io.Reader) (*http.Response, error) {
	api, _ := url.JoinPath(bc.baseUrl, uri)
	req := bc.buildRequest(ctx, http.MethodPost, api, headers, payload)
	req.Header.Set("Content-Type", "application/json")

	return bc.request(ctx, req)
}
func (bc *BaseAPI) Put(ctx context.Context, uri string, headers map[string]string, payload io.Reader) (*http.Response, error) {
	api, _ := url.JoinPath(bc.baseUrl, uri)
	req := bc.buildRequest(ctx, http.MethodPut, api, headers, payload)
	req.Header.Set("Content-Type", "application/json")

	return bc.request(ctx, req)
}
func (bc *BaseAPI) Delete(ctx context.Context, uri string, headers map[string]string, payload io.Reader) (*http.Response, error) {
	api, _ := url.JoinPath(bc.baseUrl, uri)
	req := bc.buildRequest(ctx, http.MethodDelete, api, headers, payload)
	req.Header.Set("Content-Type", "application/json")

	return bc.request(ctx, req)

}
func (bc *BaseAPI) Patch(ctx context.Context, uri string, headers map[string]string, payload io.Reader) (*http.Response, error) {
	api, _ := url.JoinPath(bc.baseUrl, uri)

	req := bc.buildRequest(ctx, http.MethodPatch, api, headers, payload)
	req.Header.Set("Content-Type", "application/json")

	return bc.request(ctx, req)
}
func (bc *BaseAPI) Head(ctx context.Context, uri string, headers map[string]string) (*http.Response, error) {
	api, _ := url.JoinPath(bc.baseUrl, uri)
	req := bc.buildRequest(ctx, http.MethodHead, api, headers, nil)
	return bc.request(ctx, req)
}
func (bc *BaseAPI) buildRequest(_ context.Context, method, uri string, headers map[string]string, payload io.Reader) *http.Request {
	uri, _ = url.QueryUnescape(uri)
	req, _ := http.NewRequest(method, uri, payload)
	req.Header.Set("Accept", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return req
}

func (bc *BaseAPI) request(ctx context.Context, req *http.Request) (*http.Response, error) {
	err := bc.requestSignature(ctx, req)
	if err != nil {
		return nil, err
	}
	return bc.cli.Do(req)
}
