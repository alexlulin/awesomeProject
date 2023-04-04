package main

import "fmt"

func vals()(int,int){
	return 3,4
}
func main() {

	a,b :=vals()
	fmt.Println(a)
	fmt.Println(b)

	//如果你仅仅需要返回值的一部分的话，你可以使用空白标识符'_'
	_,c:= vals()
	fmt.Println(c)
}
