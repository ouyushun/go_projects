package main

import "fmt"

func main() {
	res := eval(1,2, "&")
	fmt.Println(res)
}

func enums() {
	const (
		cpp = iota
		java = 5
		golang = 6
		py
	)
	fmt.Println(cpp, java, golang, py)
}

func eval(a, b int, opt string) (int, error)  {
	switch opt {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("ss")
	}
}

func pow(int)  {
	
}
