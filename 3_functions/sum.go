package main

import ("fmt")

func sum(n1 int64, n2 int64) int64 {
	return n1 + n2
}

func main() {
	var num1 int64
	var num2 int64

	fmt.Print("Enter the first number: ")
    fmt.Scan(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)
	
	the_sum := sum(num1, num2)
	fmt.Print(the_sum)

	
}