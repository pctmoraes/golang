package main

import "fmt"

// A worker pool is a concurrency pattern in which a fixed number of worker goroutines are responsible for processing tasks from a shared queue or channel. The main idea behind a worker pool is to parallelize the execution of tasks, improve efficiency, and manage resource usage by limiting the number of concurrent workers

func main() {
	tasks := make(chan int, 45)
	results := make(chan int, 45)

	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < 45; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < 45; i++ {
		resultado := <-results
		fmt.Println(resultado)
	}

}

func worker(tasks <-chan int, results chan<- int) {
	for number := range tasks {
		results <- fibonacci(number)
	}
}

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}