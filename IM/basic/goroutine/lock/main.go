package main

import (
	"fmt"
	"sync"
)

var count int
var wg sync.WaitGroup
var lock sync.Mutex

//todo 互斥锁

func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(count)
}

func add() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		lock.Lock()
		count++
		lock.Unlock()
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		lock.Lock()
		count--
		lock.Unlock()
	}
}
