package main

import "fmt"

func main() {
	// array and slice are data structs that can hold more than one value
	// like the array below that has the values of 1 to 5

	array1 := [5]int{1,2,3,4,5}
	fmt.Println(array1)

	// but the difference between an array and a slice
	// is that an array has to initialize with a stabilished size
	// like array1 = 5 elements, and this size cannot change
	// while a slice can be empty in it's initialization
	slice1 := []int{1,2}
	fmt.Println(slice1)

	// and the size can change (receive more elements)
	slice1 = append(slice1, 3)
	fmt.Println(slice1)

	// and if you look the data types with the pointer on the vars
	// aray1 is the type of [5]int and slice1 is the type of []int
	// that is a way of knowing if that var is a array or a slice
}