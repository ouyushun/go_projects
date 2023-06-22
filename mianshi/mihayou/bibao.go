package main

import (
	"fmt"
)

func foo1(x *int) func() {
	return func() {
		*x = *x + 1
		fmt.Printf("foo1 val = %d\n", *x)
	}
}



func test(x *int) {
	*x = *x + 1
	fmt.Printf("foo1 val = %d\n", *x)
}

func foo2(x int) func() {
	return func() {
		x = x + 1
		fmt.Printf("foo2 val = %d\n", x)
	}
}

func show(v interface{}) {

}
func foo4() {
	values := []int{1, 2}
	for _, val := range values {
		go func(v int) {
			fmt.Printf("foo4 val = %v\n", v)
		}(val)
	}
}
func foo5() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		go func() {
			fmt.Printf("foo5 val = %v\n", val)
		}()
	}
}
var foo6Chan = make(chan int, 10)
func foo6() {
	for val := range foo6Chan {
		go func() {
			fmt.Printf("foo6 val = %d\n", val)
		}()
	}
}
func foo7(x int) []func() {
	var fs []func()
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		fs = append(fs, func() {
			fmt.Printf("foo7 val = %d\n", x+val)
		})
	}
	return fs
}

func main() {
	// Q1第一组实验
	x := 133
	f1 := foo1(&x)
	f2 := foo2(x)

	f1()
	f1()
	f2()
	f2()
	fmt.Println("--------------------------")
	// Q1第二组
	x = 233
	f1()
	f2()
	f1()
	f2()
	// Q1第三组
	fmt.Println("main-------x = ", x)
	fmt.Println("--------------------------")
	foo1(&x)()
	foo2(x)()
	foo1(&x)()
	foo2(x)()
	foo2(x)()
}