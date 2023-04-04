package main

import (
	"fmt"
	"time"
)

func f(from string,time time.Time){
	for i:=0;i<3;i++{
		fmt.Println(from,time,":",i)
	}
}
func main() {

	//f(s) 一般会这样同步调用
	f("direct",time.Now())

	//使用‘go f(s)’ 在一个go协程中调用这个函数
	go f("goroutine",time.Now())

	go func(msg string){
		fmt.Println(msg,time.Now())
	}("going")

	fmt.Scanln()
	fmt.Println("done",time.Now())
}
