package model

import (
	"strconv"
	"strings"
)

// 字符串计算器

func GetIntByStr(str string) (reData int) {
	if len(str) == 0 {
		return
	}
	// 字符串去除空格
	str = strings.TrimSpace(str)
	// 第一步，+ 拆分
	sep := "+"
	addStrArr := getSplitArr(str, sep)
	//log.Println("--sArr",addStrArr)
	addArr := make([]string, 0) // 相加
	for j := range addStrArr {
		addStr := strings.TrimSpace(addStrArr[j])
		if isInt(addStr) {
			addArr = append(addArr, addStr)
		} else {
			// 第二步处理，
			//log.Println("--222",addStr)
			costValue := getCost(addStr)
			addArr = append(addArr, costValue)
		}

	}
	// 数组元素相加
	reData = getSum(addArr)
	//log.Println("--end ",reData)

	return
}

// 第二步处理，第一位 减后几位
func getCost(str string) (strValue string) {
	sep := "-" //减号拆分
	cStrArr := getSplitArr(str, sep)
	//log.Println("--getCost",cStrArr)
	costArr := make([]string, 0) // 相减元素
	// 第一位元素
	oneValue := ""
	for j := range cStrArr {
		costStr := strings.TrimSpace(cStrArr[j])
		if j == 0 {
			if isIntTwo(costStr) {
				oneValue = costStr
			} else {
				// 第三步处理,计算* /
				oneValue = getTake(costStr)
			}
		} else {
			if isIntTwo(costStr) {
				costArr = append(costArr, costStr)
			} else {
				// 第三步处理,计算* /
				costArr = append(costArr, getTake(costStr))
			}
		}

	}
	//  第一位减 数组和
	//log.Println("---oneValue",oneValue)
	oneValueInt, _ := strconv.Atoi(oneValue)
	sumValue := getSum(costArr)
	strValueInt := oneValueInt - sumValue
	strValue = strconv.Itoa(strValueInt)
	return
}

// 乘除计算
func getTake(str string) (strValue string) {
	sArr := make([]string, 0) // 乘除 数组
	sep := "*"                // *拆分
	cStrArr := getSplitArr(str, sep)
	//log.Println("--getTake",cStrArr)
	for j := range cStrArr {
		takeStr := strings.TrimSpace(cStrArr[j])
		if isIntThree(takeStr) {
			sArr = append(sArr, takeStr)
			sArr = append(sArr, "*")
		} else {
			sepEnd := "/" // *拆分
			endStrArr := getSplitArr(takeStr, sepEnd)
			for i := range endStrArr {
				endStr := strings.TrimSpace(endStrArr[i])
				sArr = append(sArr, endStr)
				sArr = append(sArr, "/")
			}
			sArr = append(sArr, "*")
		}
	}
	//  遍历数组，乘除
	strValue = getEnd(sArr)
	return
}

// 计算乘除
func getEnd(strArr []string) (strValue string) {
	isTake := true // 是否乘
	reValue := 1
	for j := range strArr {
		str := strArr[j]
		switch str {
		case "*":
			isTake = true
			continue
		case "/":
			isTake = false
			continue
		}
		strInt, _ := strconv.Atoi(str)
		if isTake {
			// 乘
			reValue = reValue * strInt
		} else {
			// 除
			if strInt == 0 {
				reValue = 0
			} else {
				reValue = reValue / strInt
			}
		}
	}
	strValue = strconv.Itoa(reValue)
	return
}

// 获取字符串数组和
func getSum(strArr []string) (reValue int) {
	for j := range strArr {
		str := strArr[j]
		strInt, _ := strconv.Atoi(str)
		reValue += strInt
	}
	return
}

//

func getSplitArr(str, sep string) (reData []string) {
	reData = strings.Split(str, sep)

	return
}

func isInt(str string) bool {
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

func isIntTwo(str string) bool {
	if strings.Contains(str, "*") {
		return false
	}
	if strings.Contains(str, "/") {
		return false
	}
	return true
}

func isIntThree(str string) bool {
	if strings.Contains(str, "/") {
		return false
	}
	return true
}
