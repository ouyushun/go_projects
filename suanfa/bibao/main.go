package main

import "fmt"

func main() {
	c1 := f(1)
	c2 := f(0)

	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c2())
}

func f(i int) func() int {
	return func() int {
		i++
		return i
	}
}