package main

import (
	"fmt"
	"os"
)

func main() {

	f:=createFile("d:\\cc1tt.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File{
	fmt.Println("creating")
	f,err := os.Create(p)
	if err != nil{
		panic(err)
	}
	return f
}

func writeFile(f *os.File){
	fmt.Println("writing")
	fmt.Fprintln(f,"dataadfadfafafafe你1111！！！！！！！！！----------好码")
}

func closeFile(f *os.File){
	fmt.Println("closeing")
	f.Close()
}
