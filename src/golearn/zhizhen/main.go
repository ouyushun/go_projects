package main

import "fmt"

func main() {
	a := 3
	b := 4
	swap(&a, &b)
	fmt.Println(a, b)
}

func swap(a, b *int) {
	*b, *a = *a, *b
}