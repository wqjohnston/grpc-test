package main

import (
	"context"
	"log"

	// 导入grpc包
	"google.golang.org/grpc"
	// 导入刚才我们生成的代码所在的proto包。
	pb "grpc-test/proto"
)

func main() {
	// 连接grpc服务器
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("未连接: %v", err)
	}
	// 延迟关闭连接
	defer conn.Close()

	// 建立gRPC连接
	c := pb.NewGreeterClient(conn)

	req := pb.HelloRequest{Name: "world"}

	// 调用SayHello接口，发送一条消息
	r, err := c.SayHello(context.Background(), &req)
	if err != nil {
		log.Fatalf("failed to call SayHello: %v", err)
	}

	// 打印服务的返回的消息
	log.Printf("Response: %s", r.Message)
}
