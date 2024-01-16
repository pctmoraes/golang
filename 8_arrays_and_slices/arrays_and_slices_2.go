package main

import "fmt"

func main() {
	slice1 := make([]int, 2, 3)
	fmt.Println(slice1) // output [0 0]
	
	fmt.Println(len(slice1)) // length = 2
	fmt.Println(cap(slice1)) // capacity = 3

	slice1 = append(slice1, 9)
	fmt.Println(slice1) // output [0 0 9]

	fmt.Println(len(slice1)) // length = 3
	fmt.Println(cap(slice1)) // capacity = 3

	slice1 = append(slice1, 1)
	fmt.Println(slice1) // output [0 0 9 1]

	fmt.Println(len(slice1)) // length = 4
	fmt.Println(cap(slice1)) // capacity = 6

	// what happened in lines 21 and 22?
	// the slice length and capacity changed
	// because the capacity limit were reached when in the append on line 18


	
}