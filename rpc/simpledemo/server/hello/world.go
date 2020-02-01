package hello

import (
	"log"
	"strconv"
	"time"
)

type World struct {
}

func (H *World) Say(request string, reply *string) error {
	log.Println(request)
	*reply = "Hello： " + request + ",现在是：" + time.Now().String()
	return nil
}
func (H *World) Wait(request map[string]string, reply *string) error {
	log.Println(request)
	*reply = "Hello： " + request["name"] + ",现在是：" + time.Now().String()
	return nil
}

func (H *World) Slice(request []string, reply *string) error {
	log.Println(request)
	*reply = "Hello： " + request[0] + "===" + request[1] + ",现在是：" + time.Now().String()
	//*reply = "Hello： " + request[0] + ",现在是：" + time.Now().String()
	return nil
}

func (H *World) SliceInt(request []int64, reply *string) error {
	log.Println(request)
	*reply = "Hello： " + strconv.Itoa(int(request[0])) + "===" + strconv.Itoa(int(request[1])) + ",现在是：" + time.Now().String()
	return nil
}
