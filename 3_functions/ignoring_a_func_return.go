package main

import ("fmt")


func ignoringTheFirstReturn() (string, string) {
	return "i will be ignored :(", "my name is not _ so i wont be ignored :)"
}

func ignoringTheSecondReturn() (string, string) {
	return "my name is not _ so i wont be ignored :)", "i will be ignored :("
}

func main() {
	// In order to skip the compile error 'x declared and not used'
	// we can call the return we want to ignore as _ (underscore)

	_, second_return_str := ignoringTheFirstReturn()
	fmt.Println(second_return_str)
	
	first_return_str, _ := ignoringTheSecondReturn()
	fmt.Println(first_return_str)
}