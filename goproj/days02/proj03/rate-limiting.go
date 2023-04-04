package main

import (
	"fmt"
	"time"
)

func main() {
	//首先我们将看一下基本的速率限制，假设我们想限制我们
	//接收请求处理，我们将这些请求发送给一个相同的通道

	requests := make(chan int,5)
	for i := 0; i <5 ; i++ {
		requests<-i
	}
	close(requests)

	//limiter 通道将每 200ms 接收一个值
	//这个是速率限制任务中的管理器
	limiter := time.Tick(time.Millisecond *200)
	//通过在每次请求前阻塞 ‘limiter' 通过的一个接收，我们限制自己每200ms 执行一次请求
	for req :=range requests{
		<-limiter
		fmt.Println("request0",req,time.Now())
	}

	burstyLimiter := make(chan time.Time,3)

	//想通过填充需要临时改变3次的值，做好准备
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t:= range time.Tick(time.Millisecond*200){
			burstyLimiter <-t
		}
	}()

	//现在模拟超过5个的接入请求，他们中刚开始的3个将由于受 ’bustrylimiter‘ 的脉冲 影响，可以快速的执行
	burstyRequests :=make(chan int,5)
	for i := 0; i < 5; i++ {
		burstyRequests <-i
	}

	close(burstyRequests)

	for req:=range burstyRequests{
		<-burstyLimiter
		fmt.Println("request1",req,time.Now())
	}
}
