package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var lock sync.RWMutex

//todo 读写锁

func Read() {
	defer wg.Done()
	lock.RLock() //如果只是读， 不产生影响， 如果读写同时发生， 则有影响
	fmt.Println("开始读")
	lock.RUnlock()
}

func Write() {
	defer wg.Done()
	lock.Lock()
	fmt.Println("开始写，，，，，")
	lock.Unlock()
}


func main() {
	wg.Add(6)
	for i := 0; i < 5; i++ {
		go Read()
	}
	go Write()
	wg.Wait()
}
