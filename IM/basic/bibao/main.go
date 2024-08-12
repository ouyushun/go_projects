package main

import "fmt"

func GetSum() func(int) int {
	sum := 0
	return func(i int) int {
		sum = sum + i
		return sum
	}
}

func main() {
	f := GetSum()
	fmt.Println(f(1))
	fmt.Println(f(1))
	fmt.Println(f(1))
	fmt.Println(f(1))

	g := GetSum()
	fmt.Println(g(1))
	fmt.Println(g(1))
	fmt.Println(g(1))
	fmt.Println(g(1))
}
