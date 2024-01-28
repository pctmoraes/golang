package main

import (
	"fmt"
	"time"
)

// goroutines are a fundamental concurrency primitive in go  
// they provide a way to execute functions concurrently, allowing 
// multiple tasks to be performed concurrently within a single program

func writeOnScreen(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}

}

func main() {
	go writeOnScreen("Hello")
	writeOnScreen("Welcome to goroutines")
}