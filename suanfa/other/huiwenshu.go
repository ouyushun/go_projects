package main

import (
	"fmt"
	"math"
)

func main() {
	res := revert(math.MaxInt)
	fmt.Println(res)
}

func isPalindrome(x int) bool {
	// 倒序后  判断是不是和原来的数字相等
	if x < 0 {
		return false
	}
	origin := x
	redirect := 0
	for x != 0 {
		redirect = redirect*10 + x%10
		x /= 10
	}
	return origin == redirect
}

func revert(x int) int {
	r := 0
	for x > 0 {
		r = r * 10 + x % 10
		x /= 10
	}
	return r
}
