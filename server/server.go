package main

import (
	"context"
	"log"
	"net"

	pb "grpc-test/proto"

	"google.golang.org/grpc"
)

// 简单的实现了GreeterService服务的结构体
type server struct {
	pb.UnimplementedGreeterServer
}

// 实现了在proto文件中定义的SayHello方法
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + req.Name}, nil
}

func main() {
	// 监听TCP端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建gRPC服务器实例
	s := grpc.NewServer()
	// 注册服务
	pb.RegisterGreeterServer(s, &server{})
	// 服务监听
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
