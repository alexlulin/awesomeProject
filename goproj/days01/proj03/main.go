package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var (
	timeout int64
	size int
	count int
	icmp *ICMP=&ICMP{
		Type: 8,
		Code: 0,
		CheckSum: 0,
		ID: 0,
		SequenceNum: 0,
	}
)

type ICMP struct {
	Type uint8
	Code uint8
	CheckSum uint16
	ID uint16
	SequenceNum uint16
}

func main(){
	getCommandArgs()
	desIp := os.Args[len(os.Args)-1]
	conn,err := net.DialTimeout("ip:icmp",desIp,time.Duration(timeout)*time.Millisecond)
	if err != nil{
		log.Fatal(err)
		return
	}
	defer conn.Close()

	data := make([]byte,size)
	fmt.Println(data)
}

func getCommandArgs(){
	flag.Int64Var(&timeout,"w",1000,"请求超时，单位毫秒")
	flag.IntVar(&size,"l",32,"请求发送缓冲区大小,单位字节")
	flag.IntVar(&count,"n",4,"发送请求数")
	flag.Parse()
}
