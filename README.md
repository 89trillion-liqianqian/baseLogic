一:gin http server 基础逻辑-计算器文件技术文档v1

1.目录结构

```
目录：
liqianqian@liqianqian ginserver % pwd
/Users/liqianqian/go/src/ginserver
项目结构分析：
liqianqian@liqianqian baseLogic % tree
.
├── README.md					//技术文档
├── controller				// http api
│   └── base.go				// api 
├── go.mod
├── go.sum
├── locust						//
│   ├── __pycache__
│   │   ├── load.cpython-37.pyc
│   │   └── locust.cpython-37.pyc
│   ├── load.py				//压测脚步
│   └── report_1626853467.8040588.html		//压测报告
├── main.go						//入口函数
├── model							//
│   └── requestModel.go		// model 模块
└── test
    └── test.go				// 单元测试

5 directories, 11 files
liqianqian@liqianqian baseLogic % 


```

2。运行

```
go run main.go  
```

3.api 文档

3.1

```
1）输入字符串，计算值
http get 
api: ip:port/getIntByStr?str=6/2*2/2-1
请求体
str=字符串
响应体
json
{
	"code":0,
	"msg":""ok,
	"data":2
}
状态码
0 ：请求成功

```



