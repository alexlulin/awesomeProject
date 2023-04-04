package main

import (
	"fmt"
	"time"
)

func main() {

	messages :=make(chan string,3)

	messages <- "buffered"
	messages <- "channel"
	messages <- "test"

	for i := 0; i < 3; i++ {
		fmt.Println(time.Now(),<-messages)
	}
}
