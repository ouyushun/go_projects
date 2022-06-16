package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 3, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)


	e := make([][]int, 10, 10)
	printSliceNew("e", e)
}


func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}


func printSliceNew(s string, x [][]int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}



func modifyArr(s1 []int) {
	s1[0] = 99
	fmt.Println(s1)
}

func TestSlice(s []int) {

}