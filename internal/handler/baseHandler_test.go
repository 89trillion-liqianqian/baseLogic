package handler

import (
	"log"
	"testing"
)

// 测试计算器
func TestGetIntByStrHandler(t *testing.T) {
	//str := "3+2*2+13"
	//str := "3+5 / 2 "
	str := "3//2 "
	//str := "10+2*36/2-2"
	value, err := GetIntByStrHandler(str)
	if err != nil {
		log.Println("err 计算失败", err)
	}
	log.Println("ok 计算", str, value)
}
