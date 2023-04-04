package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {

	s:="sha1 this string"
	h:=sha1.New()
	d:=md5.New()

	h.Write([]byte(s))
	d.Write([]byte(s))
	bs:=h.Sum(nil)

	cs:=d.Sum(nil)

	fmt.Println(s)
	fmt.Println("%x\n",bs)

	fmt.Println("%x\n",cs)
}
