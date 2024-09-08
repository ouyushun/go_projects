package main

import "fmt"

type A struct {
	a1 string
	a2 string
}

type B struct {
	A
}

func (a *A) sa() {
	fmt.Println("a")
}

func main() {
	b := B{}
	b.A.sa()
	b.sa()
}
