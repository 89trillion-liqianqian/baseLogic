package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

/**
api 测试
*/

// 发送Post请求
// url：         请求地址
// response：    请求返回的内容
func httpPost(urlStr, str string) string {

	resp, err := http.PostForm(urlStr,
		url.Values{
			"str": {str},
		})

	if err != nil {
		// handle error
		log.Println("--resp err")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		log.Println("--ReadAll err")
	}

	fmt.Println(string(body))
	return string(body)
}

func main() {
	result := ""
	//1）计算器接口
	url := fmt.Sprintf("http://127.0.0.1:8000/getIntByStr")
	//str := "6/2*2/2-1+2"
	//str := "3+5 / 2 "
	//str := "3/2 "
	str := "3+2*2+13"
	result = httpPost(url, str)
	log.Println("--结果", result)
}
