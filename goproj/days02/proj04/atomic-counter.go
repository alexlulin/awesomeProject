package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {

	//我们将使用一个无符号整型数来表示（永远是正整数） 这个计数器
	var ops1 uint64=0

	//为了模拟并发更新，我们启动 50个GO 协程 ，对计数器
	//每隔 1ms 进行一次 加一 操作

	for i := 0; i < 50; i++ {
		go func() {
			for{
				//使用'AddUint64' 来让计数器自动增加，使用
				//'&' 语法来给出’ops‘的内存地址
				atomic.AddUint64(&ops1,1)
				//允许其它go协程 的执行
				runtime.Gosched()
			}
		}()
	}
	//等待 一秒，让ops 的自动操作 执行一会
	time.Sleep(time.Second)

	//为了在计数器还在被其它go 协程更新时，安全的使用 它，
	//我们通过’LoadUint64‘ 将当前值的拷贝提取到’opsFinal‘
	//中 和上面一样，我们需要给这个函数取得内存地址’&ops‘

	opsFinal := atomic.LoadUint64(&ops1)

	fmt.Println("ops1:",opsFinal)
}
