package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	api "github.com/begonia-org/go-sdk/api/user/v1"
)

type TenantAPI struct {
	*BaseAPI
}
type RegisterTenantAPIResponse struct {
	*Response
	*api.Tenants
}
type GetTenantAPIResponse struct {
	*Response
	*api.Tenants
}
type ListTenantsAPIResponse struct {
	*Response
	*api.ListTenantsResponse
}

type PatchTenantsAPIResponse struct {
	*Response
	*api.Tenants
}

type DeleteTenantAPIResponse struct {
	*Response
	*api.DeleteTenantResponse
}
type AddTenantBusinessAPIResponse struct {
	*Response
	*api.TenantsBusiness
}
type ListTenantBusinessAPIResponse struct {
	*Response
	*api.ListTenantBusinessResponse
}
type DeleteTenantBusinessAPIResponse struct {
	*Response
	*api.DeleteTenantBusinessResponse
}
type PatchOptions func(in map[string]interface{}) map[string]interface{}

func WithPatchParams(key string, value interface{}) PatchOptions {
	return func(in map[string]interface{}) map[string]interface{} {
		in[key] = value
		// mask = append(mask, tiga.ToLowerCamel(key))
		// return mask
		return in
	}
}
func NewTenantAPI(addr, accessKey, secretKey string) *TenantAPI {
	return &TenantAPI{
		NewAPIClient(addr, accessKey, secretKey),
	}
}

func (t *TenantAPI) RegisterTenant(ctx context.Context, name, desc, email string, tags []string) (*RegisterTenantAPIResponse, error) {
	register := &api.PostTenantRequest{
		TenantName:  name,
		Description: desc,
		Email:       email,
		Tags:        tags,
	}
	payload, err := json.Marshal(register)
	if err != nil {
		return nil, err
	}

	rsp, err := t.Post(ctx, "/api/v1/tenants", nil, strings.NewReader(string(payload)))

	if err != nil {
		return nil, err
	}

	if rsp != nil {
		added := &api.Tenants{}
		resp, err := t.unmarshal(rsp, added)
		if err != nil && resp == nil {
			return nil, err
		}
		return &RegisterTenantAPIResponse{
			Response: resp,
			Tenants:  added,
		}, err
	}
	return nil, errors.New("response is nil")
}

func (t *TenantAPI) GetTenant(ctx context.Context, idOrName string) (*GetTenantAPIResponse, error) {

	rsp, err := t.Get(ctx, "/api/v1/tenants/"+idOrName, nil)
	if err != nil {
		return nil, err

	}
	if rsp != nil {
		details := &api.Tenants{}
		resp, err := t.unmarshal(rsp, details)
		if err != nil && resp == nil {
			return nil, err
		}
		return &GetTenantAPIResponse{
			Response: resp,
			Tenants:  details,
		}, err
	}

	return nil, errors.New("response is nil")
}
func (t *TenantAPI) ListTenants(ctx context.Context, page, pageSize int32, tags, status []string) (*ListTenantsAPIResponse, error) {
	values := url.Values{}
	for _, tag := range tags {
		values.Add("tags", tag)
	}
	for _, s := range status {
		values.Add("status", s)

	}

	values.Add("page", fmt.Sprintf("%d", page))
	values.Add("page_size", fmt.Sprintf("%d", pageSize))
	uri := fmt.Sprintf("%s?%s", "/api/v1/tenants", values.Encode())
	rsp, err := t.Get(ctx, uri, nil)
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		list := &api.ListTenantsResponse{}
		resp, err := t.unmarshal(rsp, list)
		if err != nil && resp == nil {
			return nil, err
		}
		return &ListTenantsAPIResponse{
			Response:            resp,
			ListTenantsResponse: list,
		}, err
	}
	return nil, errors.New("response is nil")
}

func (t *TenantAPI) PatchTenant(ctx context.Context, tid string, options ...PatchOptions) (*PatchTenantsAPIResponse, error) {
	in := make(map[string]interface{})
	// mask := make([]string, 0)
	for _, option := range options {
		in = option(in)
	}
	// in["update_mask"] = strings.Join(mask, ",")
	payload, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	rsp, err := t.Put(ctx, "/api/v1/tenants/"+tid, nil, strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		updated := &api.Tenants{}
		resp, err := t.unmarshal(rsp, updated)
		if err != nil && resp == nil {
			return nil, err
		}
		return &PatchTenantsAPIResponse{
			Response: resp,
			Tenants:  updated,
		}, err
	}
	return nil, errors.New("response is nil")
}

func (t *TenantAPI) AddTenantBusiness(ctx context.Context, tid string, bid string, plan string) (*AddTenantBusinessAPIResponse, error) {
	in := &api.AddTenantBusinessRequest{
		BusinessId: bid,
		TenantId:   tid,
		Plan:       plan,
	}
	payload, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	rsp, err := t.Post(ctx, "/api/v1/tenants/business", nil, strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		added := &api.TenantsBusiness{}
		resp, err := t.unmarshal(rsp, added)
		if err != nil && resp == nil {
			return nil, err
		}
		return &AddTenantBusinessAPIResponse{
			Response:        resp,
			TenantsBusiness: added,
		}, err
	}
	return nil, errors.New("response is nil")
}
func (t *TenantAPI) ListTenantBusiness(ctx context.Context, idOrName string, page, pageSize int32) (*ListTenantBusinessAPIResponse, error) {
	rsp, err := t.Get(ctx, fmt.Sprintf("/api/v1/tenants/business/%s?page=%d&page_size=%d", idOrName, page, pageSize), nil)
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		list := &api.ListTenantBusinessResponse{}
		resp, err := t.unmarshal(rsp, list)
		if err != nil && resp == nil {
			return nil, err
		}
		return &ListTenantBusinessAPIResponse{
			Response:                   resp,
			ListTenantBusinessResponse: list,
		}, err
	}
	return nil, errors.New("response is nil")
}

func (t *TenantAPI) DeleteTenantBusiness(ctx context.Context, tid string, bid string) (*DeleteTenantBusinessAPIResponse, error) {
	rsp, err := t.Delete(ctx, fmt.Sprintf("/api/v1/tenants/business/%s/%s", tid, bid), nil, nil)
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		deleted := &api.DeleteTenantBusinessResponse{}
		resp, err := t.unmarshal(rsp, deleted)
		if err != nil && resp == nil {
			return nil, err
		}
		return &DeleteTenantBusinessAPIResponse{
			Response:                     resp,
			DeleteTenantBusinessResponse: deleted,
		}, err
	}
	return nil, errors.New("response is nil")
}
func (t *TenantAPI) DeleteTenant(ctx context.Context, tid string) (*DeleteTenantAPIResponse, error) {
	rsp, err := t.Delete(ctx, fmt.Sprintf("/api/v1/tenants/%s", tid), nil, nil)
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		deleted := &api.DeleteTenantResponse{}
		resp, err := t.unmarshal(rsp, deleted)
		if err != nil && resp == nil {
			return nil, err
		}
		return &DeleteTenantAPIResponse{
			Response:             resp,
			DeleteTenantResponse: deleted,
		}, err
	}
	return nil, errors.New("response is nil")
}
