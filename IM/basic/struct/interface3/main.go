package main

import (
	"fmt"
)

// 模拟动物行为的接口
type IAnimal interface {
	Eat() // 描述吃的行为
}

type Cat struct {
	IAnimal
}

func (cat *Cat) Eat() {
	fmt.Println("cat eat")
}

func (cat *Cat) Run() {
	fmt.Println("cat eat")
}

func main() {
	cat := Cat{}
	var v IAnimal = Exec(&cat)
	fmt.Println(v)
}

func Exec(a IAnimal) IAnimal {
	a.Eat()
	return a
}
