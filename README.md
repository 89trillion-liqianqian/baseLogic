## 1.整体框架

基础逻辑-计算器

使用Gin框架完成一个接口：请求参数一个字符串，包含正整数、加(+)、减(-)、乘(*)、除(/)的算数表达式(括号除外)，返回表达式的计算结果。表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格 。 整数除法仅保留整数部分。

## 1.目录结构

```
目录：
liqianqian@liqianqian ginserver % pwd
/Users/liqianqian/go/src/ginserver
项目结构分析：
liqianqian@liqianqian baseLogic % tree
.
├── README.md										#技术文档
├── app
│   ├── http										#http server
│   │   └── httpServer.go
│   └── main.go										#入口
├── config
│   └── app.ini
├── go.mod
├── go.sum
├── internal
│   ├── ctrl										#控制层
│   │   └── base.go
│   ├── handler
│   │   ├── baseHandler.go
│   │   └── baseHandler_test.go										#单元测试
│   ├── model										#加载配置
│   │   └── config.go
│   ├── myerr										#错误
│   │   └── myerr.go
│   ├── router										#路由
│   │   └── router.go
│   └── service
│       └── service.go
├── locust										#压测
│   ├── __pycache__
│   │   ├── load.cpython-37.pyc
│   │   └── locust.cpython-37.pyc
│   ├── load.py
│   └── report_1626853467.8040588.html
├── test
│   └── test.go
└── 题二流程图.jpg

5 directories, 11 files
liqianqian@liqianqian baseLogic % 


```

## 3.逻辑代码分层

|    层     | 文件夹                           | 主要职责        | 调用关系                  | 其它说明     |
| :-------: | :------------------------------- | --------------- | ------------------------- | ------------ |
|  应用层   | /app/http/httpServer.go          | http 服务器启动 | 调用路由层                | 不可同层调用 |
|  路由层   | /internal/router/router.go       | 路由转发        | 被应用层调用，调用控制层  | 不可同层调用 |
|  控制层   | /internal/ctrl/base.go           | 计算器          | 被路由层调用，调用handler | 不可同层调用 |
| handler层 | /internal/handler/baseHandler.go | 处理具体业务    | 被控制层调用              | 不可同层调   |
|   model   | /internal/model                  | config配置加载  | 被控制层调用              |              |
| 压力测试  | Locust/load.py                   | 进行压力测试    | 无调用关系                | 不可同层调用 |

## 4.存储设计

无

## 5.接口设计供客户端调用的接口

5.1计算器

请求方法

http post 

接口地址：

127.0.0.1:8000/getIntByStr

请求参数：

```
{
	"str":3+2*2+13
}
```

json

请求响应

```
{
	"code":0,
	"msg":""ok,
	"data":20
}
```

响应状态码

| 状态码 | 说明     |
| ------ | -------- |
| 0      | 计算成功 |
| 1      | 计算失败 |

## 6.第三方库

gin

```
用于api服务，go web 框架
代码： github.com/gin-gonic/gin
```

## 7.如何编译执行

```
#切换主目录下
cd ./app/
#编译
go build
```

## 8.todo 

```
后续优化计算器算法
```

