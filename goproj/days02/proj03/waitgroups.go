package main

import (
	"fmt"
	"sync"
	"time"
)

type pio interface {
	runed() string
	moved() string
}

func worker2(id int,wg *sync.WaitGroup){
	fmt.Println("worker %d starting\n",id)

	//睡眠一秒钟，以此来模拟耗时的任务
	time.Sleep(time.Second)
	fmt.Println("Worker %d done\n",id)

	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	for i := 1; i <=5; i++ {
		wg.Add(1)
		go worker2(i,&wg)
	}

	wg.Wait()
}
