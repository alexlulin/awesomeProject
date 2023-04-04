package main

import "fmt"

func zeroval(ival int){
	ival=0
}

func zeroptr(iptr *int){
	*iptr=0
}

func main() {
	i:=1
	fmt.Println("inital:",i)

	zeroval(i)
	fmt.Println("zeroval",i)

	//通过 &i 语法来取得 ‘i’ 的 内存地址 ，即执行 'i ' 的指针
	zeroptr(&i)

	fmt.Println("zeroptr",i)

	fmt.Println("pointer:",&i)
}
