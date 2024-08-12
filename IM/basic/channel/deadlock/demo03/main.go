package main

import (
	"fmt"
	"time"
)

func main() {


	a := new(int)
	fmt.Printf("%T", a)

	var b int
	fmt.Printf("%T, %v", b, b)

	println()



	ch1 := make(chan int)
	go test(ch1)
	for {
		
		time.Sleep(time.Second)
		if v, ok :=  <- ch1; ok {
			fmt.Println(v, ok)
		} else {
			fmt.Println(v, ok)
			break
		}
	}
}

func test(ch chan int) {
	defer close(ch)
	for i := 0; i < 3; i++ {
		ch <- i + 1
	}
}