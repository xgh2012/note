package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"note/rpc/simpledemo/server/hello"
)

func main() {
	//注册
	err := rpc.RegisterName("HelloService", new(hello.World))
	if err != nil {
		log.Fatalln(err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		log.Println(conn)
		//go rpc.ServeConn(conn)
		go jsonrpc.ServeConn(conn)
	}
}
