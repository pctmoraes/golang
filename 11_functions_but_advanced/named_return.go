package main

import (
	"fmt"
)

func sumAndSubtract(n1, n2 int) (sum int, subtract int) {
	sum = n1 + n2
	subtract = n1 - n2
	return
}

func main() {
	sum, subtract := sumAndSubtract(5, 7)
	fmt.Printf("Sum: %d \nSubtract: %d", sum, subtract)
}