package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

//这个列子中，state 将被一个单独的Go 协程 拥有，这就能够保证数据在并行读取时不会混乱，为了对state 进行读取或者写入，其它的go协程 将发送
//一条数据到拥有的GO 协程中，然后接收对应的回复，结构体 ’readOp‘ 和’writeOp‘ 封装这些请求，并且是拥有Go 协程响应的一个方式
type readOp struct {
	key int
	resp1 chan int
}

type writeOp struct {
	key int
	val int
	resp1 chan bool
}


func main() {
	//计算执行操作的次数
	var readOps uint64 =0
	var writeOps uint64 =0

	//'reads' 和’writes‘ 通道分别将其他Go 协程用来发
	//读和写请求

	reads := make(chan *readOp)
	writes :=make(chan *writeOp)

	/*这个是拥有’state‘ 的那个go协程，和前面列子中的map 一样，不过这里是被这个状态协程私有的，这个GO协程
	* 反复响应到达的请求，先响应到达的请求，然后返回一个值到响应通道’resp‘ 来表示操作成功
	 */

	go func() {
		var state = make(map[int]int)
		for{
			select {
			case read := <-reads:
				read.resp1 <- state[read.key]
			case write :=<-writes:
				state[write.key]=write.val
				write.resp1<-true
			}
		}
	}()

	/**启动100个Go协程通过’reads‘ 通道发起过对 state 所有者 Go 协程的读取请求，每个读取请求需要构造要给’readOp‘,发送它到
	’reads‘ 通道中，并通过给定的’resp‘ 通道接收结果
	 */

	for r := 0; r < 100; r++ {
		go func() {
			for{
				read :=&readOp{
					key: rand.Intn(5),
					resp1: make(chan int)}
				reads <-read
				<-read.resp1
				atomic.AddUint64(&readOps,1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w:=0;w<10;w++{
		go func(){
			for{
				write := &writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp1:make(chan bool)}
				writes <- write
				<-write.resp1
				atomic.AddUint64(&writeOps,1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	//让go 协程们 跑 1s
	time.Sleep(time.Second)

	//最后 ，获取并报告’ops‘值

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:",readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:",writeOpsFinal)



}
