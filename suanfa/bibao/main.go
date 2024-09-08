package main

import (
	"fmt"
)

func main() {
	test()
	c1 := f(1)
	c2 := f(0)

	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c2())

	f1 := lazySum([]int{1, 2, 3, 4, 5})
	f2 := lazySum([]int{1, 2, 3, 4, 5, 6})
	println(f1())
	println(f2())
}

func f(i int) func() int {
	return func() int {
		i++
		return i
	}
}

// 闭包 返回的函数在其定义内部引用了局部变量
func lazySum(arr []int) func() int {
	s := 0
	return func() int {
		for _, it := range arr {
			s += it
		}
		return s
	}
}
