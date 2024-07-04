package main

import "fmt"

/*
& 是取地址符号 , 即取得某个变量的地址 , 如 ; &a
*是指针运算符 , 可以表示一个变量是指针类型 , 也可以表示一个指针变量所指向的存储单元 , 也就是这个地址所存储的值 .
*/

func main() {
	/*
	我们可以看到初始化一个指针变量，其值为nil，nil的值是不能直接赋值的。通过new其返回一个指向新分配的类型为int的指针，
	指针值为0xc00004c088，这个指针指向的内容的值为零（zero value）。
	同时，需要注意的是不同的指针类型零值是不同的。
	*/
	var v *int
	v = new(int)
	*v = 1000
	fmt.Println(*v)
	fmt.Println(v)
	fmt.Println(&v)
	fmt.Println(&*v)
}

func f3() {
	var a *int = new(int)
	fmt.Println(&a)
	fmt.Println(a) // 0xc000014330
	fmt.Println(*a) // 0

	c := 1
	d := &c
	fmt.Printf("d的类型是%T", d)
}

type Person struct {
	name string
	age int
}

func f4() {
	// & 是取地址符，取到Person类型对象的地址
	// 声明一个Person类型的结构体
	Bob := Person{"Bob", 20}
	fmt.Println("Bob:", Bob, " &Bob:", &Bob)
}

//指针的指针
func f5() {

}

func f6() {
	var a int
	var b *int
	a = 1000
	b = &a
	*b = 2000
	fmt.Println("a=", a)
	fmt.Println("b=", b)

}
