package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"note/rpc/grpc/grpcdemo/person/example"
)

func main() {
	//var dept []example.Department

	info := new(example.Person)
	info.Realname = *proto.String("熊高海")
	info.Age = *proto.Int64(923456789012345678)
	//info.Dept =

	data, err := proto.Marshal(info)
	fmt.Println(err)

	// 进行解码
	newInfo := &example.Person{}
	err = proto.Unmarshal(data, newInfo)
	if err != nil {
		fmt.Printf("unmarshaling error: ", err)
	}

	fmt.Println(newInfo.GetAge())
}
