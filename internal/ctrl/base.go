package ctrl

import (
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
		msg := "参数不正确"
		myerr.ResponseErr(c, msg)
		return
	}
	reData := handler.GetIntByStrHandler(str)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "字符串计算结果",
		"data": reData,
	})
	return
}
