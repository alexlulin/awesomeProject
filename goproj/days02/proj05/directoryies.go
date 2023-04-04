package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error){
	if e!=nil{
		panic(e)
	}
}

func main() {

	err := os.Mkdir("d:\\subdir",0755)
	check(err)

	//defer os.RemoveAll("d:\\subdir")

	createEmptyFile:= func(name string){
		d:=[]byte("")
		check(ioutil.WriteFile(name,d,0644))
	}

	createEmptyFile("d:\\subdir\\file1")

	//我们还可以创建一个有层级的目录，使用 ‘MkdirAll’ 函数，并包含其父目录
	//这个类似于 命令‘mkdir -p’

	err=os.MkdirAll("d:\\subdir\\parent\\child",0755)
	check(err)

	createEmptyFile("d:\\subdir\\parent\\file2")
	createEmptyFile("d:\\subdir\\parent\\file3")
	createEmptyFile("d:\\subdir\\parent\\child\\file4")

	c,err:= ioutil.ReadDir("d:\\subdir\\parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _,entry := range c{
		fmt.Println(" ",entry.Name(),entry.IsDir())
	}

	err= os.Chdir("d:\\subdir\\parent\\child")
	check(err)

	c,err=ioutil.ReadDir(".")
	check(err)


	fmt.Println("Listing subdir\\parent\\child")

	for _,entry:=range c{
		fmt.Println(" ",entry.Name(),entry.IsDir())
	}

	//cd 回到最开始的地方
	err =os.Chdir("../../")
	check(err)

	fmt.Println("Visiting subdir")
	err =filepath.Walk("subdir",visit)
}

func visit(p string,info os.FileInfo,err error) error{
	if err !=nil{
		return err
	}
	fmt.Println(" ",p,info.IsDir())
	return nil
}
