package controllers

import (
	"XinChat/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func GetUsers(c *gin.Context)  {

	time.Sleep(time.Second * 5)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	common.ReturnSuccess(c, 200, "success", "users")
}