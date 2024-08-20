package main

import (
	"net/http"
)

const (
	MaxWorker = 100 // 随便设置值
	MaxQueue  = 200 // 随便设置值
)

// 一个可以发送工作请求的缓冲 channel
var JobQueue chan Job

var iii int


type Payload struct{}

type Job struct {
	PayLoad Payload
}

// 接收请求，把任务筛入JobQueue。
func payloadHandler(w http.ResponseWriter, r *http.Request) {

	work := Job{PayLoad: Payload{}}
	JobQueue <- work
	_, _ = w.Write([]byte("操作成功"))
}


func main() {

	JobQueue = make(chan Job, MaxQueue)

	// 通过调度器创建worker，监听来自 JobQueue的任务
	d := NewDispatcher()
	d.Run()


	http.HandleFunc("/payload", func(writer http.ResponseWriter, request *http.Request) {
		job := Job{PayLoad: Payload{}}
		JobQueue <- job
		_, _ = writer.Write([]byte("操作成功"))
	})

	//	//创建监听端口
	err := http.ListenAndServe(":8099", nil)
	if err != nil {
		return
	}

}



