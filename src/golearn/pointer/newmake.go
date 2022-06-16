package main

import "fmt"

type People struct {
	name bool
	age int
}
func main() {
	//Type表示类型，new函数只接受一个函数，这个参数是一个类型
	//*Type表示指针类型，new函数返回一个指向该类型内存地址的指针
	a:=new(int)
	fmt.Printf("%T\n",a)//*int //指针类型
	fmt.Println(*a)//0 取值
	fmt.Println(a)//0xc00000a0a8

	var b *int
	*b = 100
	fmt.Println(*b)
}

func make(s []int) {
	//make也是用于内存分配，但是make只用于slice切片、map、chan的内存创建，
	//而且返回的是三个类型的本身，不是它们的指针类型
	//本章示例只有 声明了一个b类型的map

	//分配内存
	b :=make(map[string]int)
	//赋值
	b["沙河娜扎"] = 100
	fmt.Println(b)
}