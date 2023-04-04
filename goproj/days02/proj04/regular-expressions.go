package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	match,_:=regexp.MatchString("p([a-z]+)ch","apeachc")
	fmt.Println(match)

	r,_:= regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peacdh"))

	fmt.Println(r.FindString("punch peach"))

	fmt.Println(r.FindStringIndex("peach punch"))

	fmt.Println(r.FindStringSubmatch("pedcccch"))

	fmt.Println(r.FindStringSubmatchIndex("peach punch fmmch"))

	fmt.Println(r.FindAllString("peach punch pingch",-1))

	fmt.Println(r.FindAllStringIndex("peach punch pingch",-1))

	fmt.Println(r.Match([]byte("peach")))

	r=regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	in :=[]byte("a peach")
	out := r.ReplaceAllFunc(in,bytes.ToUpper)
	fmt.Println(string(out))
}
