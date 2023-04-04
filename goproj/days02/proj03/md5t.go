package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main(){
	has :=md5.New()
	has.Write([]byte("d:\\华为云.zip"))
	b:=has.Sum(nil)
	fmt.Printf("%x\n",hex.EncodeToString(b))
	fmt.Printf("%x\n",b)

	has1 :=md5.New()
	has1.Write([]byte("d:\\app_infos_x64.exe"))
	b1:=has1.Sum(nil)
	fmt.Printf("%x\n",hex.EncodeToString(b1))
	fmt.Printf("%x",b1)
}
