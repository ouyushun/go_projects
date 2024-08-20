package main

import (
	"fmt"
	"time"
)

func main() {
	go foo()
	for {
		fmt.Println("************")
		time.Sleep(time.Second)
	}
}

func foo() {
	go run()

}

func run() {
	for {
		go func() {
			fmt.Println("-------------------")
		}()
		time.Sleep(time.Second)
	}

}