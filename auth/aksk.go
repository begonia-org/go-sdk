package auth

type AccessKeyAuth struct{
	secretKey string
	accessKey string
	
}

func NewAccessKeyAuth(access string,secret string) *AccessKeyAuth {
	return &AccessKeyAuth{
		secretKey: secret,
		accessKey: access,
	}
}






