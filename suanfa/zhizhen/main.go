package main

import "fmt"

func main() {
	f2()
}

//ptr := &v    // v的类型为T
// v:代表被取地址的变量，类型为T
// ptr:用于接收地址的变量，ptr的类型就为*T，称做T的指针类型。*代表指针
func f1() {
	a1 := "aa"
	b1 := &a1
	fmt.Printf("%T", b1)
	fmt.Print("\n")

	/*
	用法归纳
	对变量进行取地址（&）操作，可以获得这个变量的指针变量
	对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值
	指针变量的值是变量的内存地址
	*/
	//指针取值
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
}

func f2() {
	var a = new(int)
	fmt.Printf("%T\n",a) // 0xc000014330
	fmt.Println(&a)
	fmt.Println(a) // 0xc000014330
	fmt.Println(*a) // 0
	// 变量的地址 指针的地址 指针的值

	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是%v\n", p)
	fmt.Printf("p的类型是%T\n", p)
}