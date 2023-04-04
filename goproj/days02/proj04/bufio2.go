package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	inputBuff:=bufio.NewReader(os.Stdin)
	for {
		userInput,err:=inputBuff.ReadString('\n')
		if err!=nil{
			panic(err)
		}
		fmt.Println("用户输入:",userInput[:len(userInput)-1])
		if userInput=="break"{
			break
		}
	}
	log.Println("exit...")
}
