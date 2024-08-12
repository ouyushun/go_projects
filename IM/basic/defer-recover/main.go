package main

import (
	"fmt"
)

//Recover 是一个Go语言的内建函数，可以让进入宕机流程中的 goroutine 恢复过来，
//recover 仅在延迟函数 defer 中有效，
//在正常的执行过程中，调用 recover 会返回 nil 并且没有其他任何效果，
//如果当前的 goroutine 陷入恐慌，调用 recover 可以捕获到 panic 的输入值，并且恢复正常的执行。
//通常来说，不应该对进入 panic 宕机的程序做任何处理，但有时，需要我们可以从宕机中恢复，
//至少我们可以在程序崩溃前，做一些操作，举个例子，当 web 服务器遇到不可预料的严重问题时，
//在崩溃前应该将所有的连接关闭，如果不做任何处理，会使得客户端一直处于等待状态，
//如果 web 服务器还在开发阶段，服务器甚至可以将异常信息反馈到客户端，帮助调试。


func main() {
	fmt.Println("c")
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("d")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
		}
		fmt.Println("e")
	}()
	test() //开始调用f
	fmt.Println("f") //这里开始下面代码不会再执行
}



func test() {
	//defer recover捕获错误
	a := 10
	b := 0
	r :=a/b
	panic(r)
}