package client

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	api "github.com/begonia-org/go-sdk/api/app/v1"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
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

	rsp, err := ap.Post(ctx, "/api/v1/app", nil, strings.NewReader(string(payload)))

	if err != nil {
		return nil, err
	}

	if rsp != nil {
		added := &api.AddAppResponse{}
		resp, err := ap.unmarshal(rsp, added)
		if err != nil {
			return nil, err
		}
		return &AddAppResponse{
			Response:       resp,
			AddAppResponse: added,
		}, nil

	}

	return nil, errors.New("add app config failed")
}

func (ap *AppAPI) GetAPP(ctx context.Context, appId string) (*AppDetailsResponse, error) {

	rsp, err := ap.Get(ctx, "/api/v1/app/"+appId, nil)
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

func (ap *AppAPI) PatchAPP(ctx context.Context, appid, name, description string, tags []string) (*AppDetailsResponse, error) {
	app := make(map[string]interface{})
	mask := &fieldmaskpb.FieldMask{Paths: make([]string, 0)}

	if name != "" {
		app["name"] = name
		mask.Paths = append(mask.Paths, "name")
	}
	if description != "" {
		app["description"] = description
		mask.Paths = append(mask.Paths, "description")
	}
	if len(tags) > 0 {
		app["tags"] = tags
		mask.Paths = append(mask.Paths, "tags")
	}
	app["update_mask"] = strings.Join(mask.Paths, ",")
	app["appid"] = appid
	payload, err := json.Marshal(app)
	if err != nil {
		return nil, err
	}
	rsp, err := ap.Patch(ctx, "/api/v1/app", nil, strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		updated := &api.Apps{}

		resp, err := ap.unmarshal(rsp, updated)
		if err != nil {
			return nil, err
		}
		return &AppDetailsResponse{
			Response: resp,
			Apps:     updated,
		}, nil
	}
	return nil, errors.New("delete app failed")
}

func (ap *AppAPI) DeleteAPP(ctx context.Context, appId string) (*DeleteAppResponse, error) {
	rsp, err := ap.Delete(ctx, "/api/v1/app/"+appId, nil, nil)
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		deleted := &api.DeleteAppResponse{}
		resp, err := ap.unmarshal(rsp, deleted)
		if err != nil {
			return nil, err
		}
		return &DeleteAppResponse{
			Response:          resp,
			DeleteAppResponse: deleted,
		}, nil
	}
	return nil, errors.New("delete app failed")
}
