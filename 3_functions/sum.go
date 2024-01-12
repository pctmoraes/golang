package main

import ("fmt")

func sum(n1, n2 int64) int64 {
	return n1 + n2
}

func sumAndEvenOrOdd(n1, n2 int64) (int64, string) {
	sum := n1 + n2

	var evenOrOdd string

	if sum % 2 == 0 {
		evenOrOdd = "even"
	} else {
		evenOrOdd = "odd"
	}

	return sum, evenOrOdd
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

	// The function below will return the sum and if the number is even or not
	// It is a modified version of sum() to 
	//understand how a func can return more than one value
	sum, evenOrOdd := sumAndEvenOrOdd(num1,num2)
	fmt.Println("The sum is:", sum, "and it is", evenOrOdd)
}