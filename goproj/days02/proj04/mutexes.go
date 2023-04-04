package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//在我们的例子中，’state‘ 是一个map
	var state = make(map[int]int)

	//这里的’mutex‘ 将同步’state‘的访问
	var mutex = &sync.Mutex{}

	var ops int64 =0

	for r:=0;r<100;r++{
		go func() {
			total :=0
			for{
				key := rand.Intn(5)
				mutex.Lock()
				total+=state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops,1)
				runtime.Gosched()
			}
		}()
	}

	for w:=0;w<10;w++{
		go func() {
			for{
				key:=rand.Intn(5)
				val:=rand.Intn(100)
				mutex.Lock()
				state[key]=val
				mutex.Unlock()
				atomic.AddInt64(&ops,1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:",opsFinal)

	//对'state' 使用一个最终的锁，显示它是如何结束的
	mutex.Lock()
	fmt.Println("state:",state)
	mutex.Unlock()
}
