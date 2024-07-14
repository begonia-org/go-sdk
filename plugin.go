package gosdk

import (
	"context"

	api "github.com/begonia-org/go-sdk/api/plugin/v1"
	"google.golang.org/grpc"
)

type Plugin interface {
	SetPriority(priority int)
	Priority() int
	Name() string
}
type LocalPlugin interface {
	Plugin
	UnaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error)
	StreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error
	StreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error)
}
type RemotePlugin interface {
	Plugin
	api.PluginServiceClient
}
type Plugins []Plugin

func (p Plugins) Len() int {
	return len(p)
}
func (p Plugins) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p Plugins) Less(i, j int) bool {
	return p[i].Priority() > p[j].Priority()
}
