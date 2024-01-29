package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(4) // counter = 4

	go func() {
		writeOnScreen("hello world")
		waitGroup.Done() // counter -1 = 3
	}()

	go func() {
		writeOnScreen("programming in Go!")
		waitGroup.Done() // counter -1 = 2
	}()

	go func() {
		writeOnScreen("GoRoutine 3!")
		waitGroup.Done() // counter -1 = 1
	}()

	go func() {
		writeOnScreen("Goroutine 4!")
		waitGroup.Done() // counter -1 = 0
	}()
	
	// when the counter equals zero the program stops waiting
	waitGroup.Wait() 
}

func writeOnScreen(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}