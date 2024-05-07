package client

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	api "github.com/begonia-org/go-sdk/api/user/v1"
)

type AuthzAPI struct {
	*BaseAPI
}
type AuthzSeedResponse struct {
	*Response
	*api.AuthLogAPIResponse
}
type LoginResponse struct {
	*Response
	*api.LoginAPIResponse
}

type LogoutResponse struct {
	*Response
	*api.LogoutAPIResponse
}
func NewAuthzAPI(addr, accessKey, secretKey string) *AuthzAPI {
	return &AuthzAPI{
		NewAPIClient(addr, accessKey, secretKey),
	}
}

// decryptAES 解密函数
func decryptAES(ciphertext string, secretKey string) (string, error) {
	cipherTextBytes, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	iv := cipherTextBytes[:aes.BlockSize] // IV 通常与块大小相等
	encrypted := cipherTextBytes[aes.BlockSize:]

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}

	mode := cipher.NewCFBDecrypter(block, iv)
	mode.XORKeyStream(encrypted, encrypted)

	return string(encrypted), nil
}
func (ap *AuthzAPI) Seed(ctx context.Context) (*AuthzSeedResponse, error) {
	token := time.Now().UnixMilli() * 1000
	rsp, err := ap.Get(ctx, "/api/v1/auth/log/"+fmt.Sprintf("%d", token), nil)
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		auth := &api.AuthLogAPIResponse{}
		resp, err := ap.unmarshal(rsp, auth)
		if err != nil {
			return nil, err
		}
		return &AuthzSeedResponse{
			Response:           resp,
			AuthLogAPIResponse: auth,
		}, nil
	}
	return nil, fmt.Errorf("auth seed response is nil")
}
func (ap *AuthzAPI) Login(ctx context.Context, account, password string, pubKey *rsa.PublicKey, keepLogin bool) (*LoginResponse, error) {
	auth := &api.UserAuth{
		Account:  account,
		Password: password,
	}
	payload, err := json.Marshal(auth)
	if err != nil {
		return nil, err
	}
	enc, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, payload)
	if err != nil {
		return nil, err
	}
	encodedData := base64.StdEncoding.EncodeToString(enc)
	rsp, err := ap.Seed(ctx)
	if err != nil {
		return nil, err
	}
	msg, err := decryptAES(rsp.AuthLogAPIResponse.Msg, rsp.AuthLogAPIResponse.Timestamp)
	if err != nil {
		return nil, err

	}
	authSeed := &api.AuthSeed{}
	decodedData,err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return nil, fmt.Errorf("decode seed failed: %w", err)
	}
	err = json.Unmarshal(decodedData, authSeed)
	if err != nil {
		return nil, fmt.Errorf("unmarshal auth seed failed: %w", err)
	}
	loginInfo := &api.LoginAPIRequest{
		Auth:        encodedData,
		Seed:        authSeed.Seed,
		IsKeepLogin: keepLogin,
	}
	payload, err = json.Marshal(loginInfo)
	if err != nil {
		return nil, fmt.Errorf("marshal login info failed: %w", err)
	}
	resp, err := ap.Post(ctx, "/api/v1/auth/login", nil, strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}
	if resp != nil {
		login := &api.LoginAPIResponse{}
		resp, err := ap.unmarshal(resp, login)
		if err != nil {
			return nil, err
		}
		return &LoginResponse{
			Response:         resp,
			LoginAPIResponse: login,
		}, nil
	}
	return nil, fmt.Errorf("login response is nil")

}

func (ap *AuthzAPI) Logout(ctx context.Context,token string) (*LogoutResponse, error) {
	rsp, err := ap.Delete(ctx, "/api/v1/auth/logout", map[string]string{
		"Authorization": "Bearer " + token,
	}, nil)
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		logout:= &api.LogoutAPIResponse{}
		resp, err := ap.unmarshal(rsp, logout)
		if err != nil {
			return nil, err
		}
		return &LogoutResponse{
			Response:         resp,
			LogoutAPIResponse: logout,

		}, nil
	}
	return nil, fmt.Errorf("logout response is nil")
}