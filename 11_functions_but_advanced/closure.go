package main

import "fmt"

// a closure is a function value that references variables from outside its body
// the closure allows you to capture and manipulate the environment in which it was created

func closure() func() {
	text := "This text is inside closure func"

	funct := func() {
		fmt.Println(text)
	}

	return funct
}

func main() {
	text := "This text is inside main func"
	fmt.Println(text)

	funct := closure()
	funct()
}