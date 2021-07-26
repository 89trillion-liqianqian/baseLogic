package router

import (
	"ginserver/internal/ctrl"
	"github.com/gin-gonic/gin"
)

// 路由管理
func Router(r *gin.Engine) {
	r.GET("/ping", ctrl.PingFunc)
	//r.GET("/getIntByStr", ctrl.GetIntByStr)
	r.POST("/getIntByStr", ctrl.GetIntByStr)
}
