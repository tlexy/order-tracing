package main

import (
	"base_util/trace"
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"protos"
	"syscall"
	"usersvc/internal/controller/grpc"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	ggrpc "google.golang.org/grpc"
)

func main() {
	// 初始化Jaeger tracer
	tracer, closer, err := trace.NewJaegerTracer("usersvc", "localhost:6831")
	if err != nil {
		g.Log().Error(context.TODO(), "NewJaegerTracer failed:", err)
		return
	}
	defer closer.Close()

	grpcServer := ggrpc.NewServer(
		ggrpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer)),
		ggrpc.StreamInterceptor(otgrpc.OpenTracingStreamServerInterceptor(tracer)),
	)
	protos.RegisterUserServiceServer(grpcServer, grpc.NewUserServer())

	listener, err := net.Listen("tcp", ":9091")
	if err != nil {
		panic(err)
	}

	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			panic(err)
		}
	}()
	g.Log().Warning(context.TODO(), "usersvc start at :9091")
	// 等等INT信号结束
	c := make(chan os.Signal, 1)

	// 将信号通道注册到需要捕获的信号上
	// 这里捕获了SIGINT（Ctrl+C）和SIGTERM（程序终止）信号
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// 阻塞等待接收信号
	s := <-c
	fmt.Printf("接收到信号: %v\n", s)
	g.Log().Warning(context.TODO(), "usersvc stop")
	grpcServer.GracefulStop()

}
