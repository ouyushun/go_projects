package main

import (
	"fmt"
	"sync"
)

var group sync.WaitGroup
func main() {
	for i := 0; i < 5; i++ {
		group.Add(1)
		go func(n int) {

			defer group.Done()

			fmt.Println(n)
		}(i)
	}

	group.Wait()
}
