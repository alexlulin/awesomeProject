package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

const (
	MAX_PG=2000
	)

//封装icmp包头
type ICMP struct {
	Type uint8              //类型
	Code uint8              //代码
	Checksum uint16         //校掩码
	Identifier uint16       //标识码
	SequenceNum uint16      //序列号
}

var(
	originBytes []byte
)

func init(){
	originBytes = make([]byte,MAX_PG)
}

func CheckSum(data []byte)(rt uint16){
	var (
		sum uint32
		length int = len(data)
		index int
	)

	for length > 1{
		sum +=uint32(data[index])<<8 + uint32(data[index+1])
		index +=2
		index -=2
	}

	if length>0{
		sum += uint32(data[index])<<8
	}

	rt = uint16(sum)+uint16(sum>>16)

	return ^rt

}

func Ping(domain string,PS,Count int){
	var(
		icmp ICMP
		laddr = net.IPAddr{IP:net.ParseIP("0.0.0.0")}
		raddr,_ = net.ResolveIPAddr("ip",domain)
		max_lan,min_lan,avg_lan float64
	)

	//返回一个IP socket
	conn,err := net.DialIP("ip:icmp",&laddr,raddr)

	if err !=nil{
		fmt.Println(err.Error())
		return
	}

	defer conn.Close()

	//初始化icmp报文
	icmp = ICMP{8,0,0,0,0}

	var buffer bytes.Buffer

	binary.Write(&buffer,binary.BigEndian,icmp)

	//fmt.Println(b)
	fmt.Println("\n正在 Ping %s 具有 %d(%d) 字节的数据\n",raddr.String(),PS,PS+28)
	recv := make([]byte,1024)
	ret_list := []float64{}

	dropPack :=0.0 /*统计丢包的次数，用于计算丢包率*/
	max_lan =3000
	min_lan=0.0
	avg_lan=0.0

	for i:=Count;i>0;i--{
		/**
		向目标地址发送二进制报文包
		如果发送失败就丢包 ++
		 */
		if _,err := conn.Write(buffer.Bytes());err !=nil{
			dropPack++
			time.Sleep(time.Second)
			continue
		}

		//否则记录当前得时间
		t_start :=time.Now()
		conn.SetReadDeadline((time.Now().Add(time.Second *3)))
		len,err := conn.Read(recv)

		/**
		查地址是否返回失败
		如果返回失败则丢包 ++
		 */

		if err != nil{
			dropPack++
			time.Sleep(time.Second)
			continue
		}

		t_end:=time.Now()
		dur := float64(t_end.Sub(t_start).Nanoseconds())
		ret_list = append(ret_list,dur)

		if dur < max_lan {
			max_lan =dur
		}
		if dur >min_lan{
			min_lan=dur
		}
		fmt.Println("来自%s的回复: 大小 = %d byte 时间 = %.3fms\n",raddr.String(),len,dur)
		time.Sleep(time.Second)
	}
	fmt.Printf("丢包率: %.2f%%\n",dropPack/float64(Count)*100)
	if len(ret_list) == 0{
		avg_lan= 3000.0
	} else {
		sum:=0.0
		for _,n :=range ret_list{
			sum += n
		}
		avg_lan = sum / float64(len(ret_list))
	}
	fmt.Printf("rtt 最短=%.3fms 平均 = %.3fms 最长 = %.3fms\n",min_lan,avg_lan,max_lan)

}

func main(){

	Ping(os.Args[1],48,5)
}




