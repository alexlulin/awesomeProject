package main

import "os"

func main() {

	panic("a.problem")

	_,err := os.Create("d:\file1111")
	for err !=nil{
		panic(err)
	}
}
