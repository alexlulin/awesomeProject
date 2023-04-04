package main

import (
	"archive/zip"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args)!=2{
		fmt.Println("lack of file")
		return
	}

	if !strings.Contains(os.Args[1],".zip"){
		fmt.Println("the file is not zip format")
		return
	}

	//打开一个zip读取器
	newZipReader,err:= zip.OpenReader(os.Args[1])
	if err !=nil{
		panic(err)
		return
	}

	//退出前关闭
	defer newZipReader.Close()

	//读取zip文件的信息
	for _,f:= range newZipReader.File{
		if f.FileInfo().IsDir(){
			fmt.Println(f.Name,"是一个目录")
		}else {
			fmt.Println(f.Name,"是一个文件")
		}
	}
}
