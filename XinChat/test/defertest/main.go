package main

import (
	"errors"
	"fmt"
)

func main() {
	var flag bool
	var a = 1
	fmt.Println("before flag", flag)
	fmt.Println("before a" , a)
	defer fmt.Println(a, flag)  //1 false , defer语句中的变量，在defer声明时就决定了。
	defer func() {
		fmt.Println("defer flag", flag)
		fmt.Println("defer a", a)
	}()

	flag = true
	a = 100
}


func Foo(n int64) (err error) {
	defer func() {
		if err != nil {
			fmt.Println("error in defer is :", err)
		}
	}()

	if n <= 0 {
		return errors.New("n <= 0")
	}

	return nil
}