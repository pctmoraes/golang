package main

import "fmt"

func function1() {
	fmt.Println("\nI am function 1")
}

func function2() {
	fmt.Println("\nI am function 2")
}

// defer statement is used to ensure that a function call is 
// performed later in a program's execution, usually for cleanup operations
// the defer statement schedules a function call to be ran
// after the surrounding function returns, but before the function completely exits

func gradingStudents(n1, n2 float32) string {
	defer fmt.Println("The student was:")
	fmt.Println("\nI am function gradingStudents")

	avg := (n1 + n2) / 2

	if avg >= 6 {
		return "approved :)"
	}
	return "disapproved :("
}

func main() {
	defer function1()
	function2()
	
	fmt.Println(gradingStudents(7, 2))
}