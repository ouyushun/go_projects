package controllers

import (
	"XinChat/common"
	"XinChat/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context)  {
	logger.Write("get user", "user")
	n1 := 1
	n2 := 0
	fmt.Println(n1/n2)
	common.ReturnSuccess(c, 200, "success", "users")
}