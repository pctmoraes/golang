package main

import "fmt"

// a variadic function is a function that can accept a variable number of arguments. 
//The syntax for defining a variadic function involves using 
// an ellipsis (...) followed by the type of the variable number of arguments. 
//This allows a function to take zero or more arguments of a specific type.

func printMyNumbers(text string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

func main() {
	printMyNumbers("Printing", 0, -9, 1, 39, 67, 5)
}