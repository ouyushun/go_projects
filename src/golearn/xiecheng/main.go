package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	chanDemo2()
}


//阻塞
func chanDemo2() {
	var c = make(chan int) //创建channel

	go func() {
		time.Sleep(time.Second * 2)
		c <- 1
		fmt.Println("aaa")
		c <- 2
		fmt.Println("bbb")
		c <- 3
		fmt.Println("ccc")
	}()


	fmt.Println("before")
	<- c
	fmt.Println("done")

}


//阻塞
func chanDemo() {
	var c = make(chan int) //创建channel

	go func() {
		for {
			time.Sleep(time.Second * 2)

		}
	}()
	go func() {
		c <- 1
		fmt.Println("aaa")
		c <- 2
		fmt.Println("bbb")
		c <- 3
		fmt.Println("ccc")
	}()

	fmt.Println("ddd")


	time.Sleep(time.Second * 10)

}

func makeChan()  {
	c := make(chan int)
	go func() {
		for {
			n, ok := <- c
			if !ok {
				break
			}
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
	c <- 3
	close(c)
	time.Sleep(time.Second)
}

func main2() {
	var a [10]int
	fmt.Println("Running in", runtime.Version())
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)

}