package main

import "fmt"

func main() {
	var user_age int16
	ticket := 50

	fmt.Print("Enter your age: ")
    fmt.Scan(&user_age)


    if user_age >= 12 {
        fmt.Printf("Your ticket costs $%d", ticket)
    } else if user_age >= 5 {
        fmt.Printf("Your ticket costs $%d", ticket / 2)
    } else {
        fmt.Printf("You have free ticket")
    }

	
	// this method of using a conditional is called if init
	// it allows you to declare and initialize a variable right within 
	// the if statement before evaluating a condition
	// in this case age, and inside the conditions more variables can be initialized
	// but all will be limited by this scope 
	// this is why in line 36 there is undefined error
	if age := user_age; age >= 12 {
        fmt.Printf("Your ticket costs $%d", ticket)
    } else if age >= 5 {
        fmt.Printf("Your ticket costs $%d", ticket / 2)
    } else {
        fmt.Printf("You have free ticket")
    }

	fmt.Println(age)
}
