package main

import (
	"fmt"
	"time"
)

type A struct {
	name string
	age int
}

func (a *A) t1() {
	fmt.Println(a.name)
}

func test() {

}


func main() {

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			fmt.Println(i)
		}
	}()

	select {

	}
	fmt.Println("aaaaaa")
}

