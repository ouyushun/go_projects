package main

import "fmt"

func main() {
	m1 := make(map[string]int)

	m1["a"] = 1
	m1["b"] = 1
	fmt.Println(m1)
	m2 :=  new(map[string]int)
	fmt.Println(m2)
}

