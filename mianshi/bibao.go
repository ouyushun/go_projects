package main

import "fmt"


func func1() (val int) {
	val = 10
	defer func() {
		val += 1
		fmt.Println(val)
	}()
	return val
}

func func2() (int) {
	val := 10
	defer func() {
		val += 1
		fmt.Println(val)
	}()
	return val
}

func main() {
	fmt.Println(func2())
}
