package main

import (
	"fmt"
	"time"
	"strconv"
)

// channels are communication and synchronization between goroutines
// they provide a way for one goroutine to send data to another goroutine
// channels are a fundamental part of Go's concurrency model
// and are designed to be safe and efficient

func main() {
	channel := make(chan string)
	
	go writeOnScreen("test", channel)

	fmt.Println("After writeOnScreen execution")

	for message := range channel {
		fmt.Println(message)
	}

	fmt.Println("End of program")
}

func writeOnScreen(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		msg := text + strconv.Itoa(i)
		channel <- msg
		time.Sleep(time.Second)
	}

	close(channel)
}