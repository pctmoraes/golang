package main

import (
	"fmt"
	"time"
)

// "generator" is often used to refer to a function that returns a channel. This pattern is sometimes called a generator because it allows you to produce a sequence of values asynchronously. By leveraging goroutines and channels, you can create a generator-like pattern for concurrent data generation.

func main() {
	channel := writeOnScreen("Hello world")

	for i := 0; i < 5; i++ {
		fmt.Println(<-channel)
	}
}

func writeOnScreen(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Received text: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel
}