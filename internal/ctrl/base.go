package ctrl

import (
	"fmt"
	"ginserver/internal/handler"
	"ginserver/internal/myerr"
	"github.com/gin-gonic/gin"
)

// ping
func PingFunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ping",
	})
	return
}

// 字符串计算
func GetIntByStr(c *gin.Context) {
	str := c.PostForm("str")
	if len(str) < 1 {
		msg := "参数为空"
		myerr.ResponseErr(c, msg)
		return
	}
	reData, err := handler.GetIntByStrHandler(str)
	if err != nil {
		msg := fmt.Sprintf("参数不正确:%s", err)
		myerr.ResponseErr(c, msg)
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "字符串计算结果",
		"data": reData,
	})
	return
}
