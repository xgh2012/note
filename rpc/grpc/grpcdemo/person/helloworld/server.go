package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "note/rpc/grpc/grpcdemo/person/example"
)

type service struct {
}

func (s *service) HelloWorld(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	conn, err := net.Listen("tcp", "127.0.0.1:10009")
	fmt.Println(conn, err)

	server := grpc.NewServer()

	pb.RegisterHelloServer(server, &service{})
	//读取请求并响应
	err = server.Serve(conn)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
