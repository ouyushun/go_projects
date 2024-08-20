package main


import (
	"fmt"
	"time"
)
//3.Worker: 从Dispatcher获取任务
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}


func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

// 开启一个 worker 循环监听或退出channel
func (w Worker) Start() {
	go func() {
		for {
			// 将当前的 worker 注册到 worker 队列中
			//通过for循环， 也相当于执行完JOB之后， 将jobchannel交还给workerpool
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				// 	真正业务的地方,模拟操作耗时
				time.Sleep(1000 * time.Millisecond)
				fmt.Printf("上传成功:%v\n", job)
			case <-w.quit:
				return
			}
		}
	}()
}

func (w Worker) stop() {
	go func() {
		w.quit <- true
	}()
}
