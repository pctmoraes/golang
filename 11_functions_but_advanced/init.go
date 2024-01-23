package main

import "fmt"

var n int

// the init function is a special function that is automatically invoked by the Go 
// runtime before main or another function call and it cannot be called explicitly
func init() {
	fmt.Println("Executing init function")
	n = 10
}

func main() {
	fmt.Println("Executing main function")
	fmt.Println(n)
}