package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 定义中间
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件111开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件111")
		// 执行函数
		c.Next()
		// 中间件执行完后续的一些事情
		status := c.Writer.Status()
		api, _ := c.Get("api")
		fmt.Println("中间件11111111111执行完毕", status)
		t2 := time.Since(t)
		fmt.Println(api, "time:", t2)
	}
}

func LimiterHandler() gin.HandlerFunc {
	var limitChan = make(chan bool, 3)
	return func(ctx *gin.Context) {
		fmt.Println("中间件limiter 开始执行了")

		fmt.Println("before len-------" , len(limitChan))
		limitChan <- true

		ctx.Next()

		<- limitChan
		fmt.Println("中间件limiter 执行完毕")
		fmt.Println("after len-------" , len(limitChan))
	}
}


func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()

	r := gin.Default()
	// 注册中间件
	r.Use(LimiterHandler())
	r.Use(MiddleWare())


	// {}为了代码规范
	{
		r.GET("/ce", func(c *gin.Context) {
			// 取值
			c.Set("api", "ce")
			req, _ := c.Get("request")
			fmt.Println("main--------------------ce  request:", req)
			// 页面接收
			c.JSON(200, gin.H{"request": req})
		})
	}
	r.Run()
}
