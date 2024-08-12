package main


import (
	"fmt"
	"sync"
)

type Singleton struct{}

var (
	instance Singleton
	once     sync.Once
)

func GetInstance() Singleton {
	once.Do(func() {
		instance = Singleton{}
	})
	return instance
}

//单例模式
//在单例模式中，我们需要确保一个结构体只被初始化一次。使用 sync.Once 可以轻松实现这一目标。
func main() {

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s := GetInstance()
			fmt.Printf("Singleton instance address: %p\n", &s)
		}()
	}

	wg.Wait()
}