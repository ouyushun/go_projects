package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input :=bufio.NewScanner(os.Stdin)
	fmt.Printf("%T", input)
}

func charTimes(str string, c string) int {
	res := 0
	for _, v := range str {
		if string(v) == c {
			res++
		}
	}
	return res
}
