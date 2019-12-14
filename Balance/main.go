package main

import (
	"fmt"
	"note/Balance/bservice"
	"time"
)

func main() {
	go get()
	go add()
	time.Sleep(10 * time.Millisecond)
	fmt.Println("main===>", bservice.AppIdBalance["xgh"])
}

func get() {
	baA := &bservice.Balances{AppId: "xgh"}
	ba := baA.GetBalance()
	fmt.Println("get===>", ba)
}
func add() {
	baB := &bservice.Balances{AppId: "xgh"}
	bb := baB.AddBalance(50)
	fmt.Println("add===>", bb)
}
