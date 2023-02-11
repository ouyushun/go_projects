package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	res := lastWordLength(string(data))
	fmt.Println(res)
}

func lastWordLength(str string) int {
	length := len(str)
	i := length - 1
	for ; i >= 0; i-- {
		if i < 0 {
			break
		}
		c := str[i]
		if string(c) == " " {
			break
		}
	}
	return length - i - 1
}
