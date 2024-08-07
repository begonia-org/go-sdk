package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	api "github.com/begonia-org/go-sdk/api/endpoint/v1"
	common "github.com/begonia-org/go-sdk/common/api/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

type AddEndpointConfigResponse struct {
	*Response
	Id string
}
type EndpointDetailsResponse struct {
	*Response
	*api.DetailsEndpointResponse
}
type EndpointPatchResponse struct {
	*Response
	*api.UpdateEndpointResponse
}
type EndpointDeleteResponse struct {
	*Response
	*api.DeleteEndpointResponse
}
type BegoniaClient interface {
}
type EndpointAPI struct {
	*BaseAPI
}
type EndpointListResponse struct {
	*Response
	*api.ListEndpointResponse
}
type EndpointInterface interface {
	GetBalance() string
	GetDescription() string
	GetDescriptorSet() []byte
	GetEndpoints() []*api.EndpointMeta
	GetName() string
	GetTags() []string
	ProtoReflect() protoreflect.Message
}
type EndpointOption func(EndpointInterface)

func setEndpointByName(e EndpointInterface, name string, value interface{}) {
	field := e.ProtoReflect().Descriptor().Fields().ByName(protoreflect.Name(name))
	if field == nil {
		log.Printf("field %s not found", name)
		return
	}
	e.ProtoReflect().Set(field, protoreflect.ValueOf(value))
}
func WithDescription(desc string) EndpointOption {
	return func(e EndpointInterface) {
		// c.Description = desc
		setEndpointByName(e, "description", desc)
	}

}
func WithTags(tags []string) EndpointOption {
	return func(e EndpointInterface) {
		// c.Tags = tags
		setEndpointByName(e, "tags", tags)
	}

}
func WithDescriptorSet(pb []byte) EndpointOption {
	return func(e EndpointInterface) {
		// c.DescriptorSet = pb
		setEndpointByName(e, "descriptor_set", pb)
	}

}
func WithBalance(balance string) EndpointOption {
	return func(e EndpointInterface) {
		setEndpointByName(e, "balance", balance)
	}
}
func WithEndpoints(endpoints []*api.EndpointMeta) EndpointOption {
	return func(e EndpointInterface) {
		setEndpointByName(e, "endpoints", endpoints)
	}
}
func WithName(name string) EndpointOption {
	return func(e EndpointInterface) {
		setEndpointByName(e, "name", name)
	}

}

func NewEndpointSrvConfig(opts ...EndpointOption) EndpointInterface {
	c := &api.EndpointSrvConfig{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}
func NewEndpointAPI(addr, accessKey, secretKey string) *EndpointAPI {
	return &EndpointAPI{
		NewAPIClient(addr, accessKey, secretKey),
	}
}

// PostEndpointConfig Post endpoint config to register endpoint
func (bc *BaseAPI) PostEndpointConfig(ctx context.Context, config *api.EndpointSrvConfig) (*AddEndpointConfigResponse, error) {
	payload, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	rsp, err := bc.Post(ctx, "/api/v1/endpoints", nil, strings.NewReader(string(payload)))

	if err != nil {
		return nil, err
	}

	if rsp != nil {
		added := &api.AddEndpointResponse{}
		resp, err := bc.unmarshal(rsp, added)
		if err != nil {
			return nil, err
		}
		return &AddEndpointConfigResponse{
			Response: resp,
			Id:       added.UniqueKey,
		}, nil

	}

	return nil, errors.New("add endpoint config failed")

}
func (bc *BaseAPI) GetEndpointDetails(ctx context.Context, endpointId string) (*EndpointDetailsResponse, error) {

	rsp, err := bc.Get(ctx, "/api/v1/endpoints/"+endpointId, nil)
	if err != nil {
		return nil, err

	}
	if rsp != nil {
		details := &api.DetailsEndpointResponse{}
		resp, err := bc.unmarshal(rsp, details)
		if err != nil {
			return nil, err
		}
		return &EndpointDetailsResponse{
			Response:                resp,
			DetailsEndpointResponse: details,
		}, nil
	}

	return nil, errors.New("get endpoint details failed")
}
func (bc *BaseAPI) unmarshal(rsp *http.Response, v interface{}) (*Response, error) {
	if rsp == nil {
		return nil, errors.New("response is nil")

	}
	defer rsp.Body.Close()
	requestId := rsp.Header.Get("x-request-id")
	// log.Printf("client get requestId:%s", requestId)
	apiRsp := &common.HttpResponse{}
	data, _ := io.ReadAll(rsp.Body)
	// log.Printf("response data:%s", string(data))
	err := protojson.Unmarshal(data, apiRsp)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal JSON to HttpResponse: %w", err)
	}
	// val := apiRsp.Data.AsMap()
	dataRsp := &Response{
		StatusCode: int(apiRsp.Code),
		RequestId:  requestId,
		Err:        errors.New(apiRsp.Message),
	}
	if status := apiRsp.Code; common.Code(status) != common.Code_OK {
		return dataRsp, errors.New(apiRsp.Message)
	}
	// updated := *&api.UpdateEndpointResponse{}
	jsonBytes, err := apiRsp.Data.MarshalJSON()
	if err != nil {
		return dataRsp, fmt.Errorf("Failed to marshal data to JSON: %w", err)
	}
	if err := protojson.Unmarshal(jsonBytes, v.(protoreflect.ProtoMessage)); err != nil {
		return dataRsp, fmt.Errorf("Failed to unmarshal JSON rsp to Endpoint: %w", err)
	}
	return dataRsp, nil

}
func (bc *BaseAPI) PatchEndpointConfig(ctx context.Context, config *api.EndpointSrvUpdateRequest) (*EndpointPatchResponse, error) {
	payload, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}
	pMap := map[string]interface{}{}
	_ = json.Unmarshal(payload, &pMap)
	mask := &fieldmaskpb.FieldMask{}
	if bData, err := json.Marshal(pMap["update_mask"]); err == nil {
		// pMap["mask"] = strings.Join(mask.Paths, ",")
		if err := json.Unmarshal(bData, mask); err != nil {
			return nil, err
		}
		pMap["update_mask"] = strings.Join(mask.Paths, ",")

		payload, _ = json.Marshal(pMap)

	}

	rsp, err := bc.Put(ctx, "/api/v1/endpoints/"+config.UniqueKey, nil, strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		updated := &api.UpdateEndpointResponse{}
		resp, err := bc.unmarshal(rsp, updated)
		if err != nil {
			return nil, err
		}
		return &EndpointPatchResponse{
			Response:               resp,
			UpdateEndpointResponse: updated,
		}, nil

	}

	return nil, errors.New("add endpoint config failed")

}

func (bc *BaseAPI) DeleteEndpointConfig(ctx context.Context, endpointId string) (*EndpointDeleteResponse, error) {
	rsp, err := bc.Delete(ctx, "/api/v1/endpoints/"+endpointId, nil, nil)
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		delRsp := &api.DeleteEndpointResponse{}
		resp, err := bc.unmarshal(rsp, delRsp)
		if err != nil {
			return nil, err
		}
		return &EndpointDeleteResponse{
			Response:               resp,
			DeleteEndpointResponse: delRsp,
		}, nil
	}
	return nil, errors.New("delete endpoint config failed")

}
func (bc *BaseAPI) List(ctx context.Context, tags []string, keys []string) (*EndpointListResponse, error) {
	values := url.Values{}
	for _, tag := range tags {
		values.Add("tags", tag)
	}
	for _, key := range keys {
		values.Add("unique_keys", key)
	}
	uri := "/api/v1/endpoints"
	if len(values) > 0 {
		uri = fmt.Sprintf("%s?%s", uri, values.Encode())
	}
	rsp, err := bc.Get(ctx, uri, nil)
	if err != nil {
		return nil, err

	}
	if rsp != nil {
		list := &api.ListEndpointResponse{}
		resp, err := bc.unmarshal(rsp, list)
		if err != nil {
			return nil, err
		}
		return &EndpointListResponse{Response: resp, ListEndpointResponse: list}, nil
	}
	return nil, errors.New("get endpoint list failed")
}
