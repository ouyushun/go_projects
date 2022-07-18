package main

import "fmt"

func main() {
	fmt.Println(isValid("]"))
}

func isValid(s string) bool {
	stack := []string{}
	pairs := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	isLeft := func(l string) bool {
		for _, v := range pairs {
			if l == v {
				return true
			}
		}
		return false
	}
	for _, char := range s {
		c := string(char)

		if isLeft(c) {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack) - 1]
			if pairs[c] == top {
				stack = stack[: len(stack) - 1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

