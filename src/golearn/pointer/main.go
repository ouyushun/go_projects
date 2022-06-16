package main

import (
	"fmt"
)

type tree struct {
	value int
	name string
}

func main() {
	//指针取值
	a := 1
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	fmt.Printf("value of b:%v\n", b)
	fmt.Println(a == *b)


	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)


	var ptr *int = &c
	*ptr = 91
	fmt.Println(ptr)
	fmt.Println(c)
}

