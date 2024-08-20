package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LimiterHandler() gin.HandlerFunc {
	//闭包
	//bug 当连接报错中断， 无法从channel取出造成阻塞---defer处理
	var limitChan = make(chan bool, 300)

	var i = 0
	return func(ctx *gin.Context) {
		i++
		fmt.Println("i = ", i)
		limitChan <- true
		//防止程序中途退出异常

		defer func() {
			<- limitChan
		}()
		ctx.Next()
	}
}
