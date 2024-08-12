package main

import "fmt"

func main() {
	var a []int
	a = make([]int, 1)
	a = append(a, 1,2,3)
	fmt.Println(a)

	fmt.Println(a[3])
}

func test(a interface{}) {
	if v, ok := a.(map[int]int); ok {
		v[1] = 100
	}

}