package main

import (
	"crawler.com/oys/learngo/engine"
	"crawler.com/oys/learngo/scheduler"
	"crawler.com/oys/learngo/zhenai/parser"
)

func main() {
	/*engine.SimpleEngine{}.Run(engine.Request{
		Url: "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})*/

	e := engine.ConCurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkerCount: 1}
	e.Run(engine.Request{
		Url: "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}
