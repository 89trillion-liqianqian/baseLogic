package main

import (
	"ginserver/app/http"
	"ginserver/internal/model"
)

func main() {
	// 加载配置
	filepath := "../config/app.ini"
	model.GetAppIni(filepath)
	http.HttpServer()
}
