package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		hostname,_:=os.Hostname()
		localip,_ := net.LookupIP(hostname)
		for t:= range ticker.C{
			fmt.Println("Tick at",t,hostname,localip)
		}
	}()

	time.Sleep(time.Millisecond * 2600)
	ticker.Stop()
	fmt.Println("Ticker stoped")

}
