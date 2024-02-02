package main

import (
	"fmt"
	"math/rand"
	"time"
)

// A multiplexer, often referred to as a "fan-out, fan-in" pattern, involves distributing work among multiple goroutines (fan-out) and then collecting the results from those goroutines into a single channel (fan-in). This pattern is useful for parallelizing tasks and aggregating the results. In the example below, I'll illustrate a simple fan-out, fan-in pattern using a multiplexer.

func main() {
	channel := multiplexer(writeOnScreen("Hey Ho!"), writeOnScreen("let's Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplexer(entryChannel1, entryChannel2 <-chan string) <-chan string {
	exitChannel := make(chan string)

	go func() {
		for {
			select {
			case message := <-entryChannel1:
				exitChannel <- message
			case message := <-entryChannel2:
				exitChannel <- message
			}
		}
	}()

	return exitChannel

}

func writeOnScreen(texto string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Received text: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}