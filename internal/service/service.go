package service

import (
	"strconv"
	"strings"
)

// 获取字符串数组和
func GetSum(strArr []string) (reValue int) {
	for j := range strArr {
		str := strArr[j]
		strInt, _ := strconv.Atoi(str)
		reValue += strInt
	}
	return
}

//拆分字符串
func GetSplitArr(str, sep string) (reData []string) {
	reData = strings.Split(str, sep)
	return
}

// 校验字符串
func IsInt(str string) bool {
	if strings.Contains(str, "-") {
		return false
	}
	if strings.Contains(str, "*") {
		return false
	}
	if strings.Contains(str, "/") {
		return false
	}
	return true
}

// 校验字符串
func IsIntTwo(str string) bool {
	if strings.Contains(str, "*") {
		return false
	}
	if strings.Contains(str, "/") {
		return false
	}
	return true
}

// 校验字符串
func IsIntThree(str string) bool {
	if strings.Contains(str, "/") {
		return false
	}
	return true
}
