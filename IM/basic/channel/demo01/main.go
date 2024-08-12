package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	intChan := make(chan int)

	wg.Add(2)

	go receive(intChan, &wg)
	go send(intChan, &wg)

	wg.Wait()
}

func send(ch chan int, group *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1/2)
		ch <- i
	}
	close(ch)
	group.Done()
}

func receive(ch chan int, group *sync.WaitGroup) {
	for i := range ch {
		fmt.Println(i)
	}
	for {
		i := <- ch

		fmt.Println(i)
	}


	group.Done()
}