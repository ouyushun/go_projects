package main

import "fmt"

func main() {
	_new()
}

func _make() {
	var ints []int = make([]int, 3, 6)
	fmt.Println(ints)
	ss := append(ints, 345)
	fmt.Println(ss)

	var s2 []int
	fmt.Println(s2)
}

func _new() {
	ps := new([]string)
	fmt.Println(&ps)
	fmt.Println(*ps)

}