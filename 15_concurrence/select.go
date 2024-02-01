package main

import (
	"fmt"
	"time"
)

// the select statement is used to wait on multiple communication operations. It is often used in conjunction with channels to synchronize and coordinate communication between goroutines. The select statement allows a goroutine to wait for one of multiple communication operations to complete, and it chooses the first operation that is ready.

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "Channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Channel 2"
		}
	}()

	for {
		select {
		case messagechannel1 := <-channel1:
			fmt.Println(messagechannel1)
		case messagechannel2 := <-channel2:
			fmt.Println(messagechannel2)
		}
	}

}