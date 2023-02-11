package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[1] = 1
	m[2] = 0
	delete(m, 2)
	a := make([]int, 0)
	a = append(a, 1)
	if v , ok:= m[0]; ! ok {
		fmt.Println(ok)
	} else {
		fmt.Println(v)
	}
}
