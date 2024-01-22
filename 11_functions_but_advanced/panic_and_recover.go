package main

import "fmt"

// the panic function is used to cause a runtime panic when this happens, 
// the normal flow of the program is interrupted 
// and the program starts unwinding the stack and runs the funcs that contains defer

func gradingStudents(n1, n2 float32) string {
	// comment the line below to check how panic will behave
	defer recoverExecution()

	avg := (n1 + n2) / 2

	if avg > 6 {
		return "approved :)"
	} else if avg < 6 {
		return "disapproved :("
	}

	panic("THE AVG IS EXACTLY 6")
}

// the recover function is used to regain control after a panic
// it can only be used inside a deferred function
func recoverExecution() {
	if r := recover(); r != nil {
		fmt.Println("The execution was recovered successfully")
	}
}

func main() {
	fmt.Println(gradingStudents(6, 6))
	fmt.Println("Post-execution")
}