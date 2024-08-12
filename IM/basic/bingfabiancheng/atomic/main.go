package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total uint64

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for i = 0; i < 100; i++ {
		atomic.AddUint64(&total, 1)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
//用互斥锁来保护一个数值型的共享资源，麻烦且效率低下。
//标准库的 sync/atomic 包对原子操作提供了丰富的支持。我们可以重新实现上面的例子：
	go worker(&wg)
	go worker(&wg)

	wg.Wait()
	fmt.Println(total)
}
