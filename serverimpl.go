package main

import (
	"context"
	"fmt"
	. "github.com/invictus555/greeting/api/greeting"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type GreetingServerImpl struct {
}

func getTimeString() string {
	now := time.Now()
	return now.Format("2006-01-02 15:04:05.000000000")
}

// 定义服务接口的SayHello的实现方法
func (s GreetingServerImpl) SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
	resp := new(HelloResponse)
	fmt.Printf("Get remote call from client, the context is: %s\n\n", ctx)

	resp.Message = "server(" + version + ") Response " + in.Name + " at " + getTimeString()
	fmt.Printf("Response message from server: " + resp.Message)

	return resp, nil
}

// 定义服务接口的SayBye的实现方法
func (s GreetingServerImpl) SayBye(ctx context.Context, in *emptypb.Empty) (*ByeResponse, error) {
	resp := new(ByeResponse)
	fmt.Printf("Get remote call from client, the context is: %s\n\n", ctx)

	resp.Message = "server(" + version + ") say GoodBye at" + getTimeString()

	return resp, nil
}
