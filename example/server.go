package example

import (
	"context"
	"fmt"
	"io"
	"net"

	v1 "github.com/begonia-org/go-sdk/api/example/v1"
	"google.golang.org/grpc"
)

type server struct {
	v1.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	fmt.Printf("Received: %v\n", in.GetMsg())
	return &v1.HelloReply{Message: "Hello " + in.Msg}, nil
}
func (s *server) SayHelloStreamReply(in *v1.HelloRequest, stream v1.Greeter_SayHelloStreamReplyServer) error {
	words := []string{
		"你好",
		"hello",
		"こんにちは",
		"안녕하세요",
	}

	for _, word := range words {
		data := &v1.HelloReply{
			Message: word + ":" + in.GetMsg(),
		}
		// time.Sleep(time.Second)
		// 使用Send方法返回多个数据
		if err := stream.Send(data); err != nil {
			return err
		}
	}
	return io.EOF
}
func (s *server) SayHelloGet(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	return &v1.HelloReply{Message: "Hello, world!"}, nil
}
func (s *server) SayHelloStreamSend(stream v1.Greeter_SayHelloStreamSendServer) error {
	reply := "你好:"
	for {
		// 接收客户端发来的流式数据
		res, err := stream.Recv()
		if err == io.EOF {
			// 最终统一回复
			return stream.SendAndClose(&v1.HelloReply{
				Message: reply,
			})
		}
		if err != nil {
			return err
		}

		reply += res.GetMsg()
	}
}

func (s *server) SayHelloBidiStream(stream v1.Greeter_SayHelloBidiStreamServer) error {
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
		// 返回流式响应
		if err := stream.Send(&v1.HelloReply{Message: reply}); err != nil {
			return err
		}
	}
}
func Run(addr string) {
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()                  // 创建gRPC服务器
	v1.RegisterGreeterServer(s, &server{}) // 在gRPC服务端注册服务
	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
