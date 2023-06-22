package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "LVIII"

	s = strings.ReplaceAll(s, "IV", "a")
	s = strings.ReplaceAll(s, "IX", "b")
	s = strings.ReplaceAll(s, "XL", "c")
	s = strings.ReplaceAll(s, "XC", "d")
	s = strings.ReplaceAll(s, "CD", "e")
	s = strings.ReplaceAll(s, "CM", "f")

	fmt.Println(s)

	res := 0
	for _, item := range s {
		res += getValue(string(item))
	}
	fmt.Println(res)
}



func getValue(str string) int {
	m := map[string]int{
		"I":1,
		"V":5,
		"X":10,
		"L":50,
		"C":100,
		"D":500,
		"M":1000,
		"a":4,
		"b":9,
		"c":40,
		"d":90,
		"e":400,
		"f":900,
	}
	return m[str]
}
