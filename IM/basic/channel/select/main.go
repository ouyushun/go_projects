package main

import (
	"fmt"
	"time"
)

func main() {
	test2()
}

func test2() {
	live := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			live <- true
			time.Sleep(time.Second * 1)
		}
	}()


	for {
		time.Sleep(time.Second * 2)
		select {

		case <-live:
			fmt.Println("live")

		case <- time.After(time.Second * 3):
			fmt.Println("done")
			return
		}
	}
}

func test1() {

	var c1 = make(chan int)
	var c2 = make(chan int)
	var c3 = make(chan int)


	go func() {
		for{
			time.Sleep(time.Second)
			c1 <- 100
		}
	}()

	go func() {
		for{
			time.Sleep(time.Second)
			<- c2
		}
	}()

	go func() {
		for{
			time.Sleep(time.Second)
			c3 <- 300
		}
	}()


	for {
		println("-------")
		select {
		case  <-c1:
			fmt.Println("c1")
		case c2 <- 200 :
			fmt.Println("c2")
		case <-c3:
			fmt.Println("c3")
		}
	}
}