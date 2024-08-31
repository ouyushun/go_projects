package main

import "fmt"

type Animal interface {
	Eat()
	ChangeName()
}

type Cat struct {
	Name string
}

func (cat Cat) Eat()  {
	fmt.Println("eat")
}

func (cat *Cat) ChangeName()  {
	cat.Name = "miaomiaomiaoooooo"
}

func main() {
	cat := &Cat{Name: "miaomiao"}
	cat.Eat()
	cat.ChangeName()
	fmt.Println(cat)
}
