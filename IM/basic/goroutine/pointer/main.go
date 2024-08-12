package main

import (
	"fmt"
	"time"
)

type User struct {
	name string
	age int
}

var MAP = make(map[int]*User)


func main() {
	fmt.Println(MAP)
	go func() {
		u := &User{
			"a", 18,
		}
		MAP[1] = u
	}()

	time.Sleep(time.Second * 2)

	for _, v := range MAP {
		fmt.Println(*v)
	}


	fmt.Println(MAP)


}

