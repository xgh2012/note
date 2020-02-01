package hello

import "time"

type World struct {
}

func (H *World) Say(request string, reply *string) error {
	*reply = "Hello： " + request + ",现在是：" + time.Now().String()
	return nil
}
func (H *World) Wait(request map[string]string, reply *string) error {
	*reply = "Hello： " + request["name"] + ",现在是：" + time.Now().String()
	return nil
}
