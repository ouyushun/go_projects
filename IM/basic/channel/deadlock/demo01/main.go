package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-ch)
	}()
	ch <- "send"

	wg.Wait()
}
