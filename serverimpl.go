package main

import (
	"context"
	"fmt"
	. "github.com/invictus555/greeting/api/greeting"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GreetingServerImpl struct {
}

// 定义服务接口的SayHello的实现方法
func (s GreetingServerImpl) SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
	resp := new(HelloResponse)
	fmt.Printf("Get remote call from client, the context is: %s\n\n", ctx)

	resp.Message = "rpc-server: response " + in.Name + " with " + greetingWords
	fmt.Printf("Response message from server: " + resp.Message)

	return resp, nil
}

// 定义服务接口的SayBye的实现方法
func (s GreetingServerImpl) SayBye(ctx context.Context, in *emptypb.Empty) (*ByeResponse, error) {
	resp := new(ByeResponse)
	fmt.Printf("Get remote call from client, the context is: %s\n\n", ctx)
	resp.Message = "GoodBye " + greetingWords + "..."

	return resp, nil
}
