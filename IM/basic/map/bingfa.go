package main

import (
	cmap "github.com/orcaman/concurrent-map"
	"sync"
)



var MUTEX = &sync.RWMutex{}

func testMap() {
	wg := sync.WaitGroup{}

	var _ cmap.ConcurrentMap

	var m = make(map[string]int)
	m["a"] = 1

	//1000个读
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			_ = m["a"]
		}()
	}

	//1000个并发写
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			m["a"] = 100
		}()
	}
	wg.Wait()
}


// 写入数据到map
func write(m map[string]int) {

}

// 从map中读取数据
func read(m map[string]int) {
	for {
		_ = m["a"]
	}

}
