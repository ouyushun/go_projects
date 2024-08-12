package main

import "fmt"

func main() {
	a := "看看"
	fmt.Println([]byte(a))
	fmt.Println(string([]byte{126, 230, 157, 142, 230, 172 ,163}))

	test1()
}


func test1() {
	var a int
	test11(&a)
	fmt.Println(a)
}

func test11(num *int) {
	*num = *num + 1
}