package ctrl

import (
	"ginserver/internal/handler"
	"ginserver/internal/myerr"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
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
	if strings.Contains(str, "//") || strings.Contains(str, "**") || strings.Contains(str, "++") || strings.Contains(str, "--") {
		msg := "参数不正确"
		myerr.ResponseErr(c, msg)
		return
	}
	//
	isOk := CheckStr(str)
	if !isOk {
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

// 校验字符串值合法性
func CheckStr(str string) bool {
	strArr := []rune(str)
	lenStr := len(strArr)
	for k := 0; k < lenStr; k++ {
		strV := string(strArr[k])
		if strV == "+" || strV == "-" || strV == "*" || strV == "/" {
			continue
		}
		_, err := strconv.Atoi(strV)
		if err != nil {
			return false
		}
	}
	return true
}
