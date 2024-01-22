package main

import "fmt"

// a recursive function is a function that calls itself during its execution
// instead of using a loop to repeatedly execute a set of statements
// a recursive function divides the problem into smaller subproblems and solves 
// each subproblem by calling itself, the process continues until it hits a breaking point
// and the function stops calling itself and starts returning values

func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}

func main() {

	position := uint(12)

	for i := uint(1); i <= position; i++ {
		fmt.Println(fibonacci(i))
	}
}