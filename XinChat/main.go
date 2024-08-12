package main

import (
	"XinChat/bootstrap"
	controllers "XinChat/controllers/test"
	"XinChat/global"
	"XinChat/middleware"
	"XinChat/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
// 定义中间
func myTime(c *gin.Context) {
	start := time.Now()

	// 统计时间
	since := time.Since(start)
	fmt.Println("程序用时：", since)

	c.Next()
}
func main() {
	// 设置 release模式
	//gin.SetMode(gin.ReleaseMode)
	// 或者 设置debug模式
	//gin.SetMode(gin.TestMode)

	// 初始化配置
	bootstrap.InitializeConfig()

	engine := gin.New()
	//计时器
	engine.Use(myTime)

	//日志 中间件
	engine.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	engine.Use(logger.Recover)

	//限流
	engine.Use(middleware.LimiterHandler())

	user := engine.Group("/user")
	{
		user.GET("/users", controllers.GetUsers)
	}

	engine.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})


	engine.POST("/user/add", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	engine.Run(":" + global.App.Config.App.Port)
}
