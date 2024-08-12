package main

import (
	"fmt"
)

func main() {
	test1()
}

/*
切片是一种引用类型，它有三个属性：指针，长度和容量。
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
*/
func test1() {
	var arr = [6]int{1,2,3,4,5,6}
	var s []int = arr[1:6]

	s = append(s, 7)

	fmt.Println(s)
	fmt.Println("切片的容量", cap(s))

	/*
	切片扩容的规则
	• 如果扩容之后，还没有触及原数组的容量，则切片中的指针指向的还是原数组，如果扩容后超过了原数组的容量，则开辟一块新的内存，把原来的值拷贝过来，这种情况丝毫不会影响到原数组。
	*/
	s[1] = 55
	fmt.Println("切片改变原始数组", arr)
}

