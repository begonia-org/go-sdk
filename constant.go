package gosdk

import (
	"strings"
)

const (
	MetadataKeyPrefix = "x-begonia"
)

func GetMetadataKey(key string) string {
	return MetadataKeyPrefix + "-" + strings.ToLower(key)
}

func GetHttpHeaderKey(key string) string {
	key = strings.ToLower(key)
	key = strings.TrimPrefix(key, strings.ToLower("Grpc-Metadata-"))
	if strings.HasPrefix(key, MetadataKeyPrefix) {
		return strings.TrimPrefix(key, MetadataKeyPrefix+"-")
	}
	return ""
}
