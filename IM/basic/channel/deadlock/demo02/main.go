package main

import "fmt"

func main() {
	//只在一个goroutine中读写没有缓冲区的channel会产生死锁
	//这个死锁可以通过给channel设置一个非0的缓冲区或者将读写放在不同的goroutine解决：
	var c1 chan string = make(chan string)
	c1 <- "haha"
	msg := <-c1
	fmt.Println(msg)

	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		ch1 <- "ch1 value"//无缓冲， 阻塞造成死锁
		ch2 <- "ch2 value"
	}()

	fmt.Println(<- ch2)//无缓冲， 阻塞造成死锁
	fmt.Println(<- ch1)//

}
