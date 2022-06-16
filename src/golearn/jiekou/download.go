package main

import (
	"./infra"
	"fmt"
)

func getRetiever() retriever {
	return infra.Retriever{}
}

type retriever interface {
	Get(url string) string
	Post(url string) string
}


func download(r retriever) string {
	return r.Post("http://www.imooc.com")
}

func main() {
	str := download(getRetiever())
	fmt.Println("%s\n", str)
}
