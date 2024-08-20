package main

//2. Dispatcher调度器：循环读取JobQueue
//一个Dispatcher 管理 多个Worker。

type Dispatcher struct {
	WorkerPool chan chan Job
}

func NewDispatcher() *Dispatcher {
	pool := make(chan chan Job, MaxWorker)
	return &Dispatcher{WorkerPool: pool}
}

func (d *Dispatcher) Run() {
	// 1. 开始运行 n 个 worker
	for i := 0; i < MaxWorker; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}
	// 2. 循环读取JobQueue，随机选取一个Worker执行任务
	go d.dispatch()

}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <- JobQueue:
			go func(job Job) {
				// 阻塞直到获取一个可用的worker job channel
				jobChannel := <- d.WorkerPool
				// 分发任务到 worker job channel 中
				jobChannel <- job
			}(job)
		}
	}
}
