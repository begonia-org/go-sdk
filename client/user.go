package client


import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	api "github.com/begonia-org/go-sdk/api/user/v1"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

type UsersAPI struct {
	*BaseAPI
}
type AddUserResponse struct {
	*Response
	*api.Users
}
type UsersDetailsResponse struct {
	*Response
	*api.Users
}
type DeleteUsersResponse struct {
	*Response
	*api.DeleteUserResponse
}



func NewUsersAPI(addr, accessKey, secretKey string) *UsersAPI {
	return &UsersAPI{
		NewAPIClient(addr, accessKey, secretKey),
	}
}

func (ap *UsersAPI) PostUser(ctx context.Context, user *api.PostUserRequest) (*AddUserResponse, error) {
	payload, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	rsp, err := ap.Post(ctx, "/api/v1/users", nil, strings.NewReader(string(payload)))

	if err != nil {
		return nil, err
	}

	if rsp != nil {
		added := &api.Users{}
		resp, err := ap.unmarshal(rsp, added)
		if err != nil {
			return nil, err
		}
		return &AddUserResponse{
			Response:       resp,
			Users: added,
		}, nil

	}

	return nil, errors.New("add app config failed")
}

func (ap *UsersAPI) GetUser(ctx context.Context, uid string) (*UsersDetailsResponse, error) {

	rsp, err := ap.Get(ctx, "/api/v1/users/"+uid, nil)
	if err != nil {
		return nil, err

	}
	if rsp != nil {
		details := &api.Users{}
		resp, err := ap.unmarshal(rsp, details)
		if err != nil {
			return nil, err
		}
		return &UsersDetailsResponse{
			Response: resp,
			Users:     details,
		}, nil
	}

	return nil, errors.New("get app details failed")
}

func (ap *UsersAPI) UpdateUser(ctx context.Context,uid string,updateUser map[string]interface{}) (*UsersDetailsResponse, error) {
	user := make(map[string]interface{})
	mask := &fieldmaskpb.FieldMask{Paths: make([]string, 0)}
	for k, v := range updateUser {
		user[k] = v
		mask.Paths = append(mask.Paths, k)
	
	}
	user["update_mask"] = strings.Join(mask.Paths, ",")
	payload, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	rsp, err := ap.Put(ctx, "/api/v1/users/"+uid, nil, strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		updated := &api.Users{}

		resp, err := ap.unmarshal(rsp, updated)
		if err != nil {
			return nil, err
		}
		return &UsersDetailsResponse{
			Response: resp,
			Users:     updated,
		}, nil
	}
	return nil, errors.New("delete app failed")
}

func (ap *UsersAPI) DeleteUser(ctx context.Context, uid string) (*DeleteUsersResponse, error) {
	rsp, err := ap.Delete(ctx, "/api/v1/users/"+uid, nil, nil)
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		deleted := &api.DeleteUserResponse{}
		resp, err := ap.unmarshal(rsp, deleted)
		if err != nil {
			return nil, err
		}
		return &DeleteUsersResponse{
			Response:          resp,
			DeleteUserResponse: deleted,
		}, nil
	}
	return nil, errors.New("delete app failed")
}
