package example

import (
	"context"
	"fmt"
	"net"

	api "github.com/begonia-org/go-sdk/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PluginService struct {
	api.UnimplementedPluginServiceServer
}

func (p *PluginService) Call(ctx context.Context, in *api.PluginRequest) (*api.PluginResponse, error) {
	md := make(map[string]string)
	md["key"] = "value"
	return &api.PluginResponse{Metadata: md, NewRequest: in.Request}, nil
}

func (p *PluginService) GetPluginInfo(ctx context.Context, in *emptypb.Empty) (*api.PluginInfo, error) {
	return &api.PluginInfo{Name: "example", Version: "1234"}, nil
}

func RunPlugins(addr string) {
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()                                // 创建gRPC服务器
	api.RegisterPluginServiceServer(s, &PluginService{}) // 在gRPC服务端注册服务
	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
