package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	//应使用’join' 来构建可移植（跨操作系统）的路径
	//它接收任意数量的参数，并参照传入顺序构造一个对应层次结构的路径

	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	/*
		您应该总是使用 ‘join” 代替手动 拼接 ’/' 和’\‘
		除了可移植 ，’join‘ 还会删除 多余的分隔符 和目录，使得路径 更加 规范
	*/
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	//'DIR'和 base 可以被用于 分割路径中得 目录和文件
	//此外，split 可以一次调用返回上面两个函数的结果
	fmt.Println("DIR(p)：", filepath.Dir(p))
	fmt.Println("Base(p)：", filepath.Base(p))

	//判断路径是否为绝对路径
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	//某些文件名包含了扩展名（文件类型）
	//我们可以用“Ext” 将扩展名分割出来
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	//想获取文件名清楚扩展名后的值，请使用’strings.TrimSuffix“

	fmt.Println(strings.TrimSuffix(filename, ext))

	//Rel 寻找 basepath 与 targpath 之间的相对路径
	//如果相对路径 不存在，则返回错误
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

}

