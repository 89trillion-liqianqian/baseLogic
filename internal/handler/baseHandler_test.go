package handler

import (
	"log"
	"testing"
)

// 测试计算器
func TestGetIntByStrHandler(t *testing.T) {
	str := "3+2*2+13"
	value := GetIntByStrHandler(str)
	if value != 20 {
		log.Println("err 计算失败")
	}
	log.Println("ok 计算", str, value)
}
