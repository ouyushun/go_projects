package main

import (
	"fmt"
	"time"
)

func main() {
	m := make(map[int]int)
	for i := 0; i < 10; i++ {
		m[i] = i
	}
	for i, v := range m {
		go func() {
			i = i * i
			fmt.Println("i = ", i, "v = ", v)
		}()
	}
	time.Sleep(time.Second * 1)
}
