package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check2(e error){
	if e!=nil{
		panic(e)
	}
}

func main() {
	//开始这里展示如何写入一个字符串到一个文件

	d1:=[]byte("hello\ngo\n")
	err :=ioutil.WriteFile("d:\\dat1.txt",d1,0644)

	check2(err)

	//对于更细粒度的写入，先打开一个文件

	f,err :=os.Create("d:\\dat2.txt")
	check2(err)

	//打开文件后，习惯立即使用defer 调用文件的’close‘操作
	defer f.Close()

	//你可以写入你想想写入的字节切片
	d2 :=[]byte{115,111,109,101,10}
	n2,err := f.Write(d2)
	check2(err)
	fmt.Printf("wrote %d bytes\n",n2)

	//'writestring' 也可以用
	n3,err :=f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n",n3)

	//调用‘sync’来 将缓冲去的信息写入磁盘
	f.Sync()

	//'bufio' 提供了和我们前面看到的带缓冲的读取器一样的带缓冲的写入器
	w:= bufio.NewWriter(f)
	n4,err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n",n4)

	//使用‘flush’来确保所有缓存的操作已写入底层 写入器
	w.Flush()





}
