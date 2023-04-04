package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	//向服务端发送一个http GET 请求
	//http.GET 是创建 http.Client 对象并调用其 get 方法的快捷方式
	//它使用 http.DefaultClient 对象其默认设置
	resp,err:= http.Get("http://baidu.com")

	if err!= nil{
		panic(err)
	}

	defer resp.Body.Close()

	//打印HTTP response 状态
	fmt.Println("response status:",resp.Status)

	//打印response body 前面 5 行的内容
	scanner := bufio.NewScanner(resp.Body)
	for i:=0;scanner.Scan() && i<5;i++{
		fmt.Println(scanner.Text())
	}

	if err:= scanner.Err();err !=nil{
		panic(err)
	}
}
