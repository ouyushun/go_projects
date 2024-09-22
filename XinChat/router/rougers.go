package router

import (
	"XinChat/controllers/test"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Router() *gin.Engine {
	engine := gin.Default()

	log.Printf("aaaaaaaaa")

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
	return engine
}
