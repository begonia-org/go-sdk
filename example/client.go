package example

import (
	"context"
	"io"
	"time"

	v1 "github.com/begonia-org/go-sdk/api/v1"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	gc v1.GreeterClient
}

func NewClient(addr string) *client {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return &client{
		gc: v1.NewGreeterClient(conn),
	}
}
func (c *client) SayHello() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.gc.SayHello(ctx, &v1.HelloRequest{Msg: "begonia"})
	if err != nil {
		log.Errorf("say hello error:%v", err)
		return "", err
	}
	return r.Message, nil
	// log.Infof("say hello:%s", r.Message)
}

func (c *client) SayHelloStreamReply() (<-chan string, error) {

	stream, err := c.gc.SayHelloStreamReply(context.Background(), &v1.HelloRequest{Msg: "begonia"})
	if err != nil {
		log.Errorf("say hello stream reply resp error:%v", err)
		return nil, err
	}
	ch := make(chan string, 100)
	go func() {
		defer close(ch)
		for {
			r, err := stream.Recv()
			if err != nil && io.EOF != err {
				log.Errorf("say hello stream reply recv error:%v", err)
				break
			}
			if io.EOF == err {
				break
			}
			ch <- r.Message
			log.Infof("say hello stream reply:%s", r.Message)
		}
	}()
	return ch, nil
}

func (c *client) SayHelloStreamSend() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.gc.SayHelloStreamSend(ctx)
	if err != nil {
		log.Errorf("say hello stream send error:%v", err)
		return "", err
	}
	words := []string{
		"你好",
		"hello",
		"こんにちは",
		"안녕하세요",
	}
	for _, word := range words {
		if err := stream.Send(&v1.HelloRequest{Msg: word}); err != nil {
			log.Errorf("say hello stream send error:%v", err)
		}
	}
	r, err := stream.CloseAndRecv()
	if err != nil {
		log.Errorf("say hello stream send error:%v", err)
	}
	log.Infof("say hello stream send:%s", r.Message)
	return r.Message, nil
}

func (c *client) SayHelloBidiStream() (<-chan string, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	stream, err := c.gc.SayHelloBidiStream(context.Background())
	if err != nil {
		log.Errorf("say hello bidi stream error:%v", err)
		return nil, err
	}
	words := []string{
		"你好",
		"hello",
		"こんにちは",
		"안녕하세요",
	}
	ch := make(chan string, 10)
	go func() {
		for _, word := range words {
			if err := stream.Send(&v1.HelloRequest{Msg: word}); err != nil {
				log.Errorf("say hello bidi stream error:%v", err)
			}
		}
		err := stream.CloseSend()
		if err != nil {
			log.Errorf("say hello stream send error:%v", err)
		}
	}()
	go func() {
		defer close(ch)

		r, err := stream.Recv()
		if err != nil && io.EOF != err {
			log.Errorf("say hello bidi stream error:%v", err)
		}
		if io.EOF == err {
			return
		}
		ch <- r.Message
		log.Infof("say hello bidi stream:%s", r.Message)
	}()
	return ch, nil
}
