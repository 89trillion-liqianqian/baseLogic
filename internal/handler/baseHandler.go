package handler

import (
	"fmt"
	"ginserver/internal/model"
	"ginserver/internal/service"
	"strconv"
	"strings"
	"unicode"
)

// 字符串计算器
func GetIntByStrHandler(str string) (reData int) {
	// 字符串去除空格
	str = strings.TrimSpace(str)
	// 中缀转后缀计算方式
	// 获取后缀
	postfix := infix2ToPostfix(str)
	reData = calculate(postfix)
	//////方法二拆分
	//reData =SplitFunc(str)
	return
}

// 中缀表达式转后缀表达式
func infix2ToPostfix(exp string) string {
	stack := model.ItemStack{}
	// 运算符标示
	postfix := "" //表达式
	expLen := len(exp)
	// 遍历整个表达式
	for i := 0; i < expLen; i++ {

		char := string(exp[i])

		switch char {
		case " ":
			continue
		case "(":
			// 左括号直接入栈
			stack.Push("(")
		case ")":
			// 右括号则弹出元素直到遇到左括号
			for !stack.IsEmpty() {
				preChar := stack.Top()
				if preChar == "(" {
					stack.Pop() // 弹出 "("
					break
				}
				postfix = postfix + "#" + preChar
				stack.Pop()
			}
			// 数字则直接输出
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			j := i
			digit := ""
			for ; j < expLen && unicode.IsDigit(rune(exp[j])); j++ {
				digit += string(exp[j])
			}
			postfix = postfix + "#" + digit
			i = j - 1 // i 向前跨越一个整数，由于执行了一步多余的 j++，需要减 1
		default:
			// 操作符：遇到高优先级的运算符，不断弹出，直到遇见更低优先级运算符
			for !stack.IsEmpty() {
				top := stack.Top()
				if top == "(" || isLower(top, char) {
					break
				}
				postfix = postfix + "#" + top
				stack.Pop()
			}
			// 低优先级的运算符入栈
			stack.Push(char)
		}
	}

	// 栈不空则全部输出
	for !stack.IsEmpty() {
		postfix = postfix + "#" + stack.Pop()
	}

	return postfix
}

//2、封装一个后缀表达式计算值的方法
func calculate(postfix string) int {
	stack := model.ItemStack{}
	postfix = strings.TrimSpace(postfix)
	postfixArray := strings.Split(postfix, "#")
	//fixLen := len(postfixArray)

	//fmt.Println(len(postfixArray))

	for i := 1; i < len(postfixArray); i++ {
		nextChar := postfixArray[i]
		// 数字：直接压栈
		_, err := strconv.Atoi(postfixArray[i])
		if err == nil {
			//fmt.Println("333333",postfixArray[i])
			stack.Push(postfixArray[i])
		} else {
			//fmt.Println("33333355555",postfixArray[i])
			// 操作符：取出两个数字计算值，再将结果压栈
			num1, _ := strconv.Atoi(stack.Pop())
			num2, _ := strconv.Atoi(stack.Pop())

			switch nextChar {
			case "+":
				stack.Push(strconv.Itoa(num1 + num2))
			case "-":
				stack.Push(strconv.Itoa(num1 - num2))
			case "*":
				stack.Push(strconv.Itoa(num1 * num2))
			case "/":

				stack.Push(strconv.Itoa(num2 / num1))
				fmt.Println("num:num2", num1, num2, strconv.Itoa(num1/num2))
			}
		}
	}
	result, _ := strconv.Atoi(stack.Top())
	return result
}

// 比较运算符栈栈顶 top 和新运算符 newTop 的优先级高低
func isLower(top string, newTop string) bool {
	// 优先级校验
	switch top {
	case "+", "-":
		if newTop == "*" || newTop == "/" {
			return true
		}
	case "(":
		return true
	}
	return false
}

// 拆分法
func SplitFunc(str string) (reData int) {
	// 第一步，+ 拆分
	sep := "+"
	addStrArr := service.GetSplitArr(str, sep)
	//log.Println("--sArr",addStrArr)
	addArr := make([]string, 0) // 相加
	for j := range addStrArr {
		addStr := strings.TrimSpace(addStrArr[j])
		if service.IsInt(addStr) {
			addArr = append(addArr, addStr)
		} else {
			// 第二步处理，
			costValue := getCost(addStr)
			addArr = append(addArr, costValue)
		}

	}
	// 数组元素相加
	reData = service.GetSum(addArr)
	return
}

// 第二步处理，第一位 减后几位
func getCost(str string) (strValue string) {
	sep := "-" //减号拆分
	cStrArr := service.GetSplitArr(str, sep)
	//log.Println("--getCost",cStrArr)
	costArr := make([]string, 0) // 相减元素
	// 第一位元素
	oneValue := ""
	for j := range cStrArr {
		costStr := strings.TrimSpace(cStrArr[j])
		if j == 0 {
			if service.IsIntTwo(costStr) {
				oneValue = costStr
			} else {
				// 第三步处理,计算* /
				oneValue = getTake(costStr)
			}
		} else {
			if service.IsIntTwo(costStr) {
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
	sumValue := service.GetSum(costArr)
	strValueInt := oneValueInt - sumValue
	strValue = strconv.Itoa(strValueInt)
	return
}

// 乘除计算
func getTake(str string) (strValue string) {
	sArr := make([]string, 0) // 乘除 数组
	sep := "*"                // *拆分
	cStrArr := service.GetSplitArr(str, sep)
	//log.Println("--getTake",cStrArr)
	for j := range cStrArr {
		takeStr := strings.TrimSpace(cStrArr[j])
		if service.IsIntThree(takeStr) {
			sArr = append(sArr, takeStr)
			sArr = append(sArr, "*")
		} else {
			sepEnd := "/" // *拆分
			endStrArr := service.GetSplitArr(takeStr, sepEnd)
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
