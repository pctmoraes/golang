package main 

import "fmt"

func main() { 
	// an anonymous function is a function that is defined without a name
	// instead of declaring a separate function with a name, 
	//you can define a function on the fly and use it where it's needed.	
	
	// AF with no parameter and no return 
	func(){ 
		fmt.Println("hello! i am an anonymous func!") 
	}()


	// AF with a string as a parameter and a return 
	my_return := func(text string) string { 
		return "returning: " + text
	}("my_text")

	fmt.Println(my_return) 

	


	
} 
