package main

import (
	"flag"
	"fmt"
	. "github.com/invictus555/greeting/api/greeting"
	"google.golang.org/grpc"
	"net"
	"os"
)

var (
	port          *int
	greetingWords string
)

var greetingServerImpl = GreetingServerImpl{}

func init() {
	// 获取环境变量GREETING_WORDS
	greetingWords = os.Getenv("GREETING_WORDS")
	if greetingWords == "" {
		greetingWords = "default"
	}

	// 设置获取GreetingServer监听端口的命令行参数名称，默认值，使用说明
	port = flag.Int("port", 50050, "the port of rpc server")
	flag.Parse() // 解析命令行参数
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()                         // 实现gRPC Server
	RegisterGreetingServer(s, greetingServerImpl) // 注册helloServer为客户端提供服务, 内部调用了s.RegisterServer()
	println("Listen address and port " + fmt.Sprintf(":%d", *port))

	_ = s.Serve(listen)
}
