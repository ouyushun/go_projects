package main

import (
	"fmt"
	"runtime"
	"time"
)

func SayHello() {
	for {
		fmt.Println("Hello goroutinue")
	}
}

func main() {
	defer func() {
		fmt.Println("the number of goroutine is: ", runtime.NumGoroutine())
	}()
	go SayHello()
	fmt.Println("Hello main")
	time.Sleep(time.Second * 1)
}


