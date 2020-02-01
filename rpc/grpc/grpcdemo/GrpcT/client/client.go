package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "note/rpc/grpc/grpcdemo/GrpcT/helloWorld"
	"os"
	"time"
)

func main() {
	//连接服务端句柄
	//WithInsecure()不指定安全选项
	conn, err := grpc.Dial("127.0.0.1:50001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	cli := pb.NewHelloClient(conn)

	name := "World XGH"
	//获取命令行参数
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	//设置上下文超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//响应
	resp, err := cli.HelloWorld(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not succ: %v", err)
	}
	log.Printf("Receive from server: %s", resp.Message)
}
