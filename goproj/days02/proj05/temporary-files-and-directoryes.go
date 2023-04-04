package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check2(e error){
	if e!=nil{
		panic(e)
	}
}

func main() {

	//创建临时文件最简单的方法是调用‘ioutil.TempFile’ 函数
	//它会创建并打开文件，我们可以对文件进行读写
	//函数的第一个参数将‘“”’ ，‘ioutil.TempFile’ 会操作系统的默认位置下创建该文件
	f,err := ioutil.TempFile("","sample")
	check2(err)
	//打印临时文件
	fmt.Println("temp file name",f.Name())

	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1,2,3,4})
	check2(err)

	dname,err:=ioutil.TempDir("","sampledir")
	fmt.Println("temp dir name",dname)

	defer os.RemoveAll(dname)

	fname:= filepath.Join(dname,"file1")
	err= ioutil.WriteFile(fname,[]byte{1,2},0666)
	check2(err)

}


