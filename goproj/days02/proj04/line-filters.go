package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//对‘os.stdin '使用 一个带缓冲的 scanner
	//让我们可以直接使用方便’scan‘ 方法来直接读取一行
	//每次调用该方法可以让scanner 读取下一行

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan(){
		//'text' 返回当前的 token,现在是输入的下一行
		ucl := strings.ToUpper(scanner.Text())
		//写出大写的行
		fmt.Println(ucl)
	}

	//检查’scan‘的错误，文件结束符是可以接收的，并且不会被 scan 当作 一个错误

	if err:= scanner.Err();err !=nil{
		fmt.Fprintln(os.Stderr,"error：",err)
		os.Exit(1)
	}

}
