package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

var (
	ch            = make(chan bool)
	limits        = rate.NewLimiter(1, 2)
	success, errs int
	isSusscee     bool
)

func main() {
	for j := 0; j < 100; j++ {
		for i := 0; i < 100; i++ {
			go func() {
				inlist()
			}()
		}
		time.Sleep(1000 * time.Millisecond)
		out()
	}
	fmt.Println("main ====> ", success, errs)
}

func out() {
	for i := 0; i < 100; i++ {
		isSusscee = <-ch
		if isSusscee == true {
			success++
		} else if isSusscee == false {
			errs++
		}
	}
	fmt.Println(success, errs)
}

func inlist() {
	if limits.Allow() {
		ch <- true
		return
	}
	ch <- false
}
