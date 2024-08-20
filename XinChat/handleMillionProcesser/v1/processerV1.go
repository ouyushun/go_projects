package v1

import (
	"fmt"
	"net/http"
	"time"
)

// Payload https://cloud.tencent.com/developer/article/1831381
//这样操作存在什么问题呢？一般情况下，没什么问题。
//但是如果是高并发的场景下，不对 goroutine数进行控制，你的 CPU 使用率暴涨，内存占用暴涨，直至程序奔溃。
//
//如果此操作落地至数据库，例如 mysql,那么相应的，你数据库的服务器磁盘IO、网络带宽 、CPU负载、内存消耗都会达到非常高的情况，一并奔溃。所以，一旦程序中出现不可控的事物，往往是危险的信号。
type Payload struct {
	// 传啥不重要
}

func (p *Payload) UpdateToS3() error {
	//存储逻辑,模拟操作耗时
	time.Sleep(500 * time.Millisecond)
	fmt.Println("上传成功")
	return nil
}

func payloadHandler(w http.ResponseWriter, r *http.Request) {
	// 业务过滤
	// 请求body解析......
	var p Payload
	go p.UpdateToS3()
	w.Write([]byte("操作成功"))
}

