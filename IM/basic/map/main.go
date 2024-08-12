package main

import (
	"fmt"
	"time"
)

func main() {
	testMap()


	time.Sleep(time.Second * 3)
}



func test1()  {
	//创建
	//1. 只声明， 没用分配内存空间
	//var m map[int]string
	m := make(map[int]string)
	m[1] = "a"
	m[2] = "a"
	m[3] = "a"
	fmt.Println(m)
	fmt.Printf("%T", m)

	//2.
}

func test2() {
	m := make(map[int]int)
	m[1] = 1
	m[2] = 2
	m[3] = 3
	test3(m)
	fmt.Println(m)
}

func test3(m map[int]int) {
	for k, v := range m {
		m[k] = v * 2
	}
}