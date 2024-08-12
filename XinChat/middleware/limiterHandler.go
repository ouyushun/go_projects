package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)


var m = make(map[string]string)

func LimiterHandler() gin.HandlerFunc {
	//闭包
	//bug 当连接报错中断， 无法从channel取出造成阻塞
	var limitChan = make(chan bool, 100)

	var i = 0
	return func(ctx *gin.Context) {
		i++
		fmt.Println("i = ", i)
		fmt.Println("before len-------" , len(limitChan))
		limitChan <- true

		ctx.Next()

		<- limitChan
		fmt.Println("after len-------" , len(limitChan))
	}
}
