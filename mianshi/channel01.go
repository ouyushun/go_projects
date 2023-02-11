package main

import (
	"fmt"
	"time"
)

func goRead(a <-chan int) {
	val, ok := <-a
	fmt.Println("goRoutineA received the data", val, ok)
}

func goWrite(a chan <- int) {
	a <- 11
}

func main() {
	ch := make(chan int)

	//go goWrite(ch)
	go goRead(ch)

	time.Sleep(time.Second * 2)
}
