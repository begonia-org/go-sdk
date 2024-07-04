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

type BusinessAPI struct {
	*BaseAPI
}
type AddBusinessAPIResponse struct {
	*Response
	*api.Business
}
type GetBusinessAPIResponse struct {
	*Response
	*api.Business
}
type ListBusinessAPIResponse struct {
	*Response
	*api.ListBusinessResponse
}
type DeleteBusinessAPIResponse struct {
	*Response
	*api.DeleteBusinessResponse
}
type PatchBusinessAPIResponse struct {
	*Response
	*api.Business
}

func NewBusinessAPI(addr, accessKey, secretKey string) *BusinessAPI {
	return &BusinessAPI{
		NewAPIClient(addr, accessKey, secretKey),
	}
}
func (b *BusinessAPI) PostBusiness(ctx context.Context, name, desc string, tags []string) (*AddBusinessAPIResponse, error) {
	business := &api.PostBusinessRequest{
		BusinessName: name,
		Description:  desc,
		Tags:         tags,
	}
	payload, err := json.Marshal(business)
	if err != nil {
		return nil, err
	}

	rsp, err := b.Post(ctx, "/api/v1/business", nil, strings.NewReader(string(payload)))

	if err != nil {
		return nil, err
	}

	if rsp != nil {
		added := &api.Business{}
		resp, err := b.unmarshal(rsp, added)
		// if err != nil {
		// 	// log.Printf("post business error:%v", err)

		// 	return nil, err
		// }
		return &AddBusinessAPIResponse{
			Response: resp,
			Business: added,
		}, err
	}
	return nil, errors.New("response is nil")
}

func (b *BusinessAPI) GetBusiness(ctx context.Context, idOrName string) (*GetBusinessAPIResponse, error) {
	rsp, err := b.Get(ctx, "/api/v1/business/"+idOrName, nil)
	if err != nil {
		return nil, err
	}

	if rsp != nil {
		added := &api.Business{}
		resp, err := b.unmarshal(rsp, added)
		// if err != nil {
		// 	return resp, err
		// }
		return &GetBusinessAPIResponse{
			Response: resp,
			Business: added,
		}, err
	}
	return nil, errors.New("response is nil")
}
func (b *BusinessAPI) ListBusiness(ctx context.Context, tags []string, page, pageSize int32) (*ListBusinessAPIResponse, error) {
	values := url.Values{}
	for _, tag := range tags {
		values.Add("tags", tag)
	}
	values.Add("page", fmt.Sprintf("%d", page))
	values.Add("page_size", fmt.Sprintf("%d", pageSize))
	uri := fmt.Sprintf("%s?%s", "/api/v1/business", values.Encode())

	rsp, err := b.Get(ctx, uri, make(map[string]string))
	if err != nil {
		return nil, err
	}

	if rsp != nil {
		added := &api.ListBusinessResponse{}
		resp, err := b.unmarshal(rsp, added)
		// if err != nil {
		// 	return nil, err
		// }
		return &ListBusinessAPIResponse{
			Response:             resp,
			ListBusinessResponse: added,
		}, err
	}
	return nil, errors.New("response is nil")
}

func (b *BusinessAPI) DeleteBusiness(ctx context.Context, idOrName string) (*DeleteBusinessAPIResponse, error) {
	rsp, err := b.Delete(ctx, "/api/v1/business/"+idOrName, nil, nil)
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		deleted := &api.DeleteBusinessResponse{}
		resp, err := b.unmarshal(rsp, deleted)
		// if err != nil {
		// 	return nil, err
		// }
		return &DeleteBusinessAPIResponse{
			Response:               resp,
			DeleteBusinessResponse: deleted,
		}, err
	}
	return nil, errors.New("response is nil")
}
func (b *BusinessAPI) PatchBusiness(ctx context.Context, bid string, options ...PatchOptions) (*PatchBusinessAPIResponse, error) {
	bs := make(map[string]interface{})
	// mask := make([]string, 0)
	for _, option := range options {
		bs = option(bs)

	}
	// bs["update_mask"] = strings.Join(mask, ",")
	// log.Printf("patch mask:%v", mask)
	payload, err := json.Marshal(bs)
	// log.Printf("patch payload:%v", string(payload))
	if err != nil {
		return nil, err
	}

	rsp, err := b.Put(ctx, "/api/v1/business/"+bid, make(map[string]string), strings.NewReader(string(payload)))

	if err != nil {
		return nil, err
	}

	if rsp != nil {
		added := &api.Business{}
		resp, err := b.unmarshal(rsp, added)
		// if err != nil {
		// 	return nil, err
		// }
		return &PatchBusinessAPIResponse{
			Response: resp,
			Business: added,
		}, err
	}
	return nil, errors.New("response is nil")
}
