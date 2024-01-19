package main

import (
	"fmt"
)

// a named return value is a feature that allows you 
// to declare and name the return variables in the function signature
// this simplifies the code by making it clear what each return variable represents
// it also has a special impact on the return behavior

func sumAndSubtract(n1, n2 int) (sum int, subtract int) {
	sum = n1 + n2
	subtract = n1 - n2
	return
}

func main() {
	sum, subtract := sumAndSubtract(5, 7)
	fmt.Printf("Sum: %d \nSubtract: %d", sum, subtract)
}