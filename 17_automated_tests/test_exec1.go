package main

import (
	"fmt"
	"tests/address"
)

func main() {
	first_type := address.AddressType("street of the silly")
	fmt.Println(first_type)

	second_type := address.AddressType("avenue 123")
	fmt.Println(second_type)

	third_type := address.AddressType("saint paul square")
	fmt.Println(third_type)
}