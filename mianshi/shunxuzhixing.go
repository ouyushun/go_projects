package main

import (
	"fmt"
	"sync"
	"time"
)

type task struct{
	cond    *sync.Cond
	fn  func()
	t   *task
}

func (t *task)run(){
	go func(){
		t.cond.L.Lock()
		defer t.cond.L.Unlock()
		wg.Done()
		t.cond.Wait()

		t.fn()

		if t.t!=nil {

			t.t.send()
		}


	}()

}


func (t *task)send(){
	t.cond.L.Lock()
	defer t.cond.L.Unlock()
	t.cond.Signal()

}


func f1(){

	time.Sleep(time.Second*1)
	fmt.Println("f1")
}

func f2(){
	time.Sleep(time.Second*1)
	fmt.Println("f2")
}


func f3(){

	time.Sleep(time.Second*1)
	fmt.Println("f3")
}

var wg *sync.WaitGroup

func main(){

	wg = &sync.WaitGroup{}
	wg.Add(3)
	t1 := task{fn:f1}
	t2 := task{fn:f2,t:&t1}
	t3 := task{fn:f3,t:&t2}

	t1.cond = sync.NewCond(&sync.Mutex{})
	t2.cond = sync.NewCond(&sync.Mutex{})
	t3.cond = sync.NewCond(&sync.Mutex{})

	t1.run()
	t2.run()
	t3.run()

	wg.Wait()
	t3.send()

	time.Sleep(time.Second*20)
	fmt.Println("ok")
}
