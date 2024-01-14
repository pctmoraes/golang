package main

import "fmt"

func main() {
	// setting var1 with value of ten and copying var1 value to var2 
	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	// setting var3 with value of two
	// setting a pointer to var3
	var var3 int = 2
	var pointer *int = &var3

	fmt.Printf("\nvar3 value: %d | var3 memory address: %p", var3, pointer)
	fmt.Printf("\npointer's memory address value: %d\n", *pointer)

	var3 += 6
	fmt.Printf("\nvar3 value: %d | var3 memory address: %p", var3, pointer)
	fmt.Printf("\npointer's memory address value: %d", *pointer)

}