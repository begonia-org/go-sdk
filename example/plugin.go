package example

import (
	"context"
	"fmt"
	"net"
	"strings"

	api "github.com/begonia-org/go-sdk/api/plugin/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PluginService struct {
	api.UnimplementedPluginServiceServer
}

func (p *PluginService) Apply(ctx context.Context, in *api.PluginRequest) (*api.PluginResponse, error) {
	md := make(map[string]string)
	md["key"] = "value"
	if inmd, ok := metadata.FromIncomingContext(ctx); ok {
		for k, v := range inmd {
			md[k] = strings.Join(v, ",")
		}
	}
	newMd := metadata.New(md)
	err := grpc.SetHeader(ctx, newMd)
	if err != nil {
		return nil, err

	}
	return &api.PluginResponse{NewRequest: in.Request}, nil
}
func (p *PluginService) Metadata(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	md := make(map[string]string)
	md["key"] = "value"
	if inmd, ok := metadata.FromIncomingContext(ctx); ok {
		for k, v := range inmd {
			md[k] = strings.Join(v, ",")
		}
	}
	newMd := metadata.New(md)
	err := grpc.SetHeader(ctx, newMd)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
func (p *PluginService) Info(ctx context.Context, in *emptypb.Empty) (*api.PluginInfo, error) {
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
