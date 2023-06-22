package main

import "fmt"

func main() {
	fmt.Println(test("d"))
}

func test(a any) any {
	return a
}
