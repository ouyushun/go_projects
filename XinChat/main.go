package main

import (
	"XinChat/bootstrap"
	controllers "XinChat/controllers/test"
	"XinChat/global"
	"XinChat/middleware"
	"XinChat/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kr/pretty"
	"github.com/spf13/viper"
	"net/http"
)

func main() {

	// 设置 release模式
	//gin.SetMode(gin.ReleaseMode)
	// 或者 设置debug模式
	//gin.SetMode(gin.TestMode)

	// 初始化配置
	bootstrap.InitializeConfig()

	pretty.Println(viper.Get("app"))

	fmt.Println("------------")

	engine := gin.New()
	//日志 中间件  logrus  替换掉default log
	engine.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	engine.Use(logger.Recover)


	//计时器 中间件
	engine.Use(middleware.MyTime())

	//限流 中间件
	engine.Use(middleware.LimiterHandler())


	//日志 中间件  zap
	// 初始化日志
	//global.App.Log = bootstrap.InitializeLog()
	//global.App.Log.Info("log init success!")



	user := engine.Group("/user")
	{
		user.GET("/users", controllers.GetUsers)
	}

	engine.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	engine.Run(":" + global.App.Config.App.Port)
}
