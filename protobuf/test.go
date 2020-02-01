package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"note/protobuf/protofile"
)

func main() {

	real := proto.String("Xgh2012")
	age := proto.Int64(25)
	// 创建一个消息 Person
	info := &protofile.Person{
		Age:      *age,
		Realname: *real,
		Sex:      "",
		Vehicle:  nil,
		Skill:    0,
		WorkDay:  nil,
		Dept:     nil,
	}
	// 进行编码
	data, err := proto.Marshal(info)
	if err != nil {
		fmt.Printf("marshaling error: ", err)
	}

	// 进行解码
	newInfo := &protofile.Person{}
	err = proto.Unmarshal(data, newInfo)
	if err != nil {
		fmt.Printf("unmarshaling error: ", err)
	}

	fmt.Println(newInfo.GetRealname())
	fmt.Println(newInfo.GetAge())
	fmt.Println(newInfo)
}
