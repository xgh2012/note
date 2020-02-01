package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

func main() {
	//client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("client", err)
		return
	}

	var (
		res string
		t   = make(map[string]string)
	)
	t["name"] = "xgh"

	//err = client.Call("HelloService.Say","Xgh",&res)
	err = client.Call("HelloService.Wait", t, &res)
	if err != nil {
		fmt.Println("res", err)
		return
	}
	fmt.Println(res)
}
