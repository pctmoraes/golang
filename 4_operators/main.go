package main

import "fmt"

func main() {
	sum := 2 + 3
	subtraction := 2 - 3
	multiplication := 2 * 3
	division := 2 / 3
	_mod := 2 % 3

	fmt.Println(sum, subtraction, multiplication, division, _mod)

	// Above there is nothing new to prog. lang. when it comes to operators
	// but in go, the code below will generate error
	// this happens because go do not allow two different types on an operation
	var x int8 = 5
	var y int16 = 12

	xy_sum := x + y
	fmt.Println(xy_sum)

	
	// to solve this make sure to equal the vars types
	xy_sum2 := int16(x) + y
	fmt.Println((xy_sum2))
}