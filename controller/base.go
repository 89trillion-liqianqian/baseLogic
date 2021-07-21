package controller

import (
	"ginserver/model"
	"github.com/gin-gonic/gin"
)

/**
http get post 处理
*/

func Routers(r *gin.Engine) {
	r.GET("/ping", pingFunc)
	r.GET("/getIntByStr", GetIntByStr)
	return
}

// 测试
func pingFunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong22",
	})
	return
}

// 字符串计算

func GetIntByStr(c *gin.Context) {
	str := c.Query("str")
	//log.Println("--str",str)
	//str:="3+2*2+13"  // 20
	reData := model.GetIntByStr(str)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "字符串计算结果",
		"data": reData,
	})
	return
}
