package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)



func MyTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 统计时间
		since := time.Since(start)
		fmt.Println("程序用时：", since)
		fmt.Println("****************************************************")
		c.Next()
	}
}
