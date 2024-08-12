package main

import "fmt"

type Teacher struct {
	Name string
	Age int
}

func main() {
	var t1 Teacher

	t1.Age = 18
	t1.Name = "a"
	fmt.Println(t1)
}
