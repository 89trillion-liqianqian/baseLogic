package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

/**
api 测试
*/

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容

func httpGet(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

func main() {
	result := ""
	//1）计算器接口
	//aa:= fmt.Sprintf("http://127.0.0.1:8000/getIntByStr?str=%s","10-2")
	aa := fmt.Sprintf("http://127.0.0.1:8000/getIntByStr?str=%s", "6/2*2/2-1")
	//result= httpGet(ipStr+api)
	result = httpGet(aa)
	log.Println("--结果", result)
}
