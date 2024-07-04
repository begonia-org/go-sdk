package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	api "github.com/begonia-org/go-sdk/api/app/v1"
)

type AppAPI struct {
	*BaseAPI
}
type AddAppResponse struct {
	*Response
	*api.AddAppResponse
}
type AppDetailsResponse struct {
	*Response
	*api.Apps
}
type DeleteAppResponse struct {
	*Response
	*api.DeleteAppResponse
}
type ListAPPResponse struct {
	*Response
	*api.AppsListResponse
}

const (
	ADD_APP_URI = "/api/v1/app"
)

func NewAppAPI(addr, accessKey, secretKey string) *AppAPI {
	return &AppAPI{
		NewAPIClient(addr, accessKey, secretKey),
	}
}

func (ap *AppAPI) PostAppConfig(ctx context.Context, app *api.AppsRequest) (*AddAppResponse, error) {
	payload, err := json.Marshal(app)
	if err != nil {
		return nil, err
	}

	rsp, err := ap.Post(ctx, "/api/v1/apps", nil, strings.NewReader(string(payload)))

	if err != nil {
		return nil, err
	}

	if rsp != nil {
		added := &api.AddAppResponse{}
		resp, err := ap.unmarshal(rsp, added)
		if err != nil && resp==nil {
			return nil, err
		}
		return &AddAppResponse{
			Response:       resp,
			AddAppResponse: added,
		}, err

	}

	return nil, errors.New("add app config failed")
}

func (ap *AppAPI) GetAPP(ctx context.Context, appId string) (*AppDetailsResponse, error) {

	rsp, err := ap.Get(ctx, "/api/v1/apps/"+appId, nil)
	if err != nil {
		return nil, err

	}
	if rsp != nil {
		details := &api.Apps{}
		resp, err := ap.unmarshal(rsp, details)
		if err != nil {
			return nil, err
		}
		return &AppDetailsResponse{
			Response: resp,
			Apps:     details,
		}, nil
	}

	return nil, errors.New("get app details failed")
}

func (ap *AppAPI) UpdateAPP(ctx context.Context, appid string, options ...PatchOptions) (*AppDetailsResponse, error) {
	app := make(map[string]interface{})
	mask := make([]string, 0)
	for _, option := range options {
		app = option(app)

	}
	app["update_mask"] = strings.Join(mask, ",")
	app["appid"] = appid
	payload, err := json.Marshal(app)
	if err != nil {
		return nil, err
	}
	rsp, err := ap.Put(ctx, "/api/v1/apps", nil, strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		updated := &api.Apps{}

		resp, err := ap.unmarshal(rsp, updated)
		// if err != nil {
		// 	return nil, err
		// }
		return &AppDetailsResponse{
			Response: resp,
			Apps:     updated,
		}, err
	}
	return nil, errors.New("delete app failed")
}

func (ap *AppAPI) DeleteAPP(ctx context.Context, appId string) (*DeleteAppResponse, error) {
	rsp, err := ap.Delete(ctx, "/api/v1/apps/"+appId, nil, nil)
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		deleted := &api.DeleteAppResponse{}
		resp, err := ap.unmarshal(rsp, deleted)
		// if err != nil {
		// 	return nil, err
		// }
		return &DeleteAppResponse{
			Response:          resp,
			DeleteAppResponse: deleted,
		}, err
	}
	return nil, errors.New("delete app failed")
}

func (ap *AppAPI) ListAPP(ctx context.Context, tags []string, status []api.APPStatus, page, pageSize int) (*ListAPPResponse, error) {
	values := url.Values{}
	for _, tag := range tags {
		values.Add("tags", tag)
	}
	for _, st := range status {
		values.Add("status", st.String())
	}
	values.Add("page", fmt.Sprint(page))
	values.Add("page_size", fmt.Sprint(pageSize))
	rsp, err := ap.Get(ctx, "/api/v1/apps?"+values.Encode(), nil)
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		apps := &api.AppsListResponse{}
		resp, err := ap.unmarshal(rsp, apps)
		if err != nil {
			return nil, err
		}
		return &ListAPPResponse{
			Response:         resp,
			AppsListResponse: apps,
		}, nil
	}
	return nil, errors.New("list app failed")
}
