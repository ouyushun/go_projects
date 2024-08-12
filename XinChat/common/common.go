package common

import "github.com/gin-gonic/gin"

type JsonStruct struct {
	Code int `json:"code"`
	Msg interface{} `json:"msg"`
	Data interface{} `json:"data"`
}


func ReturnSuccess(c *gin.Context, Code int, Msg string, Data interface{})  {
	c.JSONP(200, JsonStruct{Code: Code, Msg: Msg, Data: Data})
}

func ReturnError(c * gin.Context, Code int, Msg string)  {
	c.JSONP(200, JsonStruct{Code: Code, Msg: Msg})
}