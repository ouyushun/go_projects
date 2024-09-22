package main

import (
	"fmt"
)

// 模拟动物行为的接口
type IAnimal interface {
	Eat() // 描述吃的行为
}

// 动物 所有动物的父类
type Animal struct {
	Name string
}

// 动物去实现IAnimal中描述的吃的接口
func (a *Animal) Eat() {
	fmt.Printf("%v is eating\n", a.Name)
}

func (a *Animal) Run() {
	fmt.Printf("%v  animal is running\n", a.Name)
}

// 动物的构造函数
func newAnimal(name string) *Animal {
	return &Animal{
		Name: name,
	}
}

// 猫的结构体 组合了animal
type Cat struct {
	*Animal
}

func (c *Cat) Run() {
	fmt.Printf("%v is running\n", c.Name)
}

// 实现猫的构造函数 初始化animal结构体
func newCat(name string) *Cat {
	return &Cat{
		Animal: newAnimal(name),
	}
}

func main() {
	cat := newCat("cat")
	cat.Animal.Eat() // cat is eating
	cat.Animal.Run() // cat is eating
	cat.Run()        // cat is eating

	var cc interface{} = cat

	v := cc.(IAnimal)
	fmt.Println(v.Eat)
}
