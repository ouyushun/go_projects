package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	__scan()
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	fmt.Println(data)
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
}


func __scan() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		 r := scanner.Bytes()
		 fmt.Println(r)
	}
}

// 填fmt.scanf的坑
func _Scanf(input *string) {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	*input = string(data)
}