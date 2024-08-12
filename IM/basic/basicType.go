package main

import (
	"fmt"
	"strconv"
)

func main() {
	IntType()
}


func IntType()  {
	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 9223372036854775807
	println(integer8, integer16, integer32, integer64)
}

func Cov()  {
	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Printf("%T, %T, %T", -42, i, s)
}