package example

import (
	"testing"

	api "github.com/begonia-org/go-sdk/api/example/v1"
	"google.golang.org/protobuf/proto"
)

func TestProto(t *testing.T) {
	reply := &api.HelloReply{Message: "ddddddeeeeee"}
	data, _ := proto.Marshal(reply)
	t.Log(string(data))
}
