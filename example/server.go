package example

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"

	v1 "github.com/begonia-org/go-sdk/api/example/v1"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc"
)

var s *grpc.Server

type server struct {
	v1.UnimplementedGreeterServer
}

func NewExampleServer() *server {
	return &server{}
}
func (s *server) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	// fmt.Printf("Received: %v\n", in.GetMsg())
	return &v1.HelloReply{Message: in.Msg, Name: in.Name}, nil
}
func (s *server) SayHelloServerSideEvent(in *v1.HelloRequest, stream v1.Greeter_SayHelloServerSideEventServer) error {
	for i := 0; i < 10; i++ {
		data := &v1.HelloReply{
			Message: fmt.Sprintf("%s-%d", in.Msg, i),
			Name:    in.Name,
		}
		if err := stream.Send(data); err != nil {
			return err
		}

	}
	return nil
}
func (s *server) SayHelloGet(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	return &v1.HelloReply{Message: in.Msg, Name: in.Name}, nil

}
func (s *server) SayHelloClientStream(stream v1.Greeter_SayHelloClientStreamServer) error {
	// reply := "你好:"、
	replies := make([]*v1.HelloReply, 0)
	index := 0
	for {
		// 接收客户端发来的流式数据
		res, err := stream.Recv()

		if err == io.EOF {
			// 最终统一回复
			return stream.SendAndClose(&v1.RepeatedReply{Replies: replies})
		}
		if err != nil {
			return err
		}
		replies = append(replies, &v1.HelloReply{
			Message: fmt.Sprintf("%s-%d", res.GetMsg(), index),
			Name:    res.GetName(),
		})
		index++

		// reply += res.GetMsg()
	}
}

func (s *server) SayHelloWebsocket(stream v1.Greeter_SayHelloWebsocketServer) error {
	index := 0

	for {
		// 接收流式请求
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		reply := in.GetMsg() // 对收到的数据做些处理
		msg := fmt.Sprintf("%s-%d", reply, index)
		index++
		// 返回流式响应
		if err := stream.Send(&v1.HelloReply{Message: msg, Name: in.GetName()}); err != nil {
			return err
		}
	}
}
func (s *server) SayHelloBody(ctx context.Context, in *httpbody.HttpBody) (*httpbody.HttpBody, error) {
	req := v1.HelloRequest{}
	err := json.Unmarshal(in.Data, &req)
	if err != nil {
		log.Printf("json.Unmarshal error: %v", err)
		return nil, err
	}
	reply := &v1.HelloReply{
		Message: req.Msg,
		Name:    req.Name,
	}
	data, err := json.Marshal(reply)
	if err != nil {
		log.Printf("json.Marshal error: %v", err)
		return nil, err

	}
	return &httpbody.HttpBody{
		Data: data,
	}, nil

}
func (s *server) Desc() *grpc.ServiceDesc {
	return &v1.Greeter_ServiceDesc
}

func Run(addr string) {
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s = grpc.NewServer()                   // 创建gRPC服务器
	v1.RegisterGreeterServer(s, &server{}) // 在gRPC服务端注册服务
	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)

		return
	}
}
func Stop() {
	if s == nil {
		return
	}
	s.GracefulStop()
	s.Stop()
}
