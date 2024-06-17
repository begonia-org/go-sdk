package gosdk

import (
	"strings"
)

const (
	MetadataKeyPrefix = "x-begonia"
)
const AccessKeyType = "access_key"
const ApiKeyType = "api_key"
const UidType = "uid"

func GetMetadataKey(key string) string {
	return MetadataKeyPrefix + "-" + strings.ToLower(key)
}

func GetHttpHeaderKey(key string) string {
	key = strings.ToLower(key)
	key = strings.TrimPrefix(key, strings.ToLower("Grpc-Metadata-"))
	if strings.HasPrefix(key, MetadataKeyPrefix) {
		return strings.TrimPrefix(key, MetadataKeyPrefix+"-")
	}
	return key
}
