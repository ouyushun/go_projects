package main

import "fmt"

func main() {
	var arr [4]int
	arr[0] = 1
	arr[1] = 1
	arr[2] = 1
	fmt.Println(arr)

	var a *int
	b := 3
	a = &b
	fmt.Println(a)
}
