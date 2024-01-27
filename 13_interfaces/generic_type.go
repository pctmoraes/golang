package main

import "fmt"

// the interface below was constructed in a generic way meaning that it can
// take any type and process without running on errors

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic("String")
	generic(1)
	generic(true)

	
	// the Println func is a example of an interface in go
	fmt.Println(1, 2, "string", false, true, float64(12345))


	// one scenario that this would not be recomended is the one below
	// where a map can accept any type in its keys and values
	map1 := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "String",
		true:         float64(12),
	}

	fmt.Println(map1)
}