package main

import (
	"crawler.com/oys/learngo/engine"
	"crawler.com/oys/learngo/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url: "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}
