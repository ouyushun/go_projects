package main

import (

	"fmt"
)

type Cat struct {
	name string
	age int
}

func (c Cat) say() string {
	return c.name
}


type Dog struct {
	name string
	age int
}

func (d *Dog) say() string {
	d.name = "new_doge"
	return d.name
}

//接口（interface）是一种类型
func main() {
	var x  Animal

	var c = Cat{name: "cat", age: 1}
	var d = Dog{name: "dog", age: 1}

	x = c
	fmt.Println(x.say())

	x = &d
	//x = d
	fmt.Println(d.say())
	fmt.Println(d)

	fmt.Printf("%T,%T", x, d)
}
