package main

import "fmt"

func main() {
	// a map is a data-structer of unordered collections of key-value pairs
	// like a dictionary in other programming languages
	// the keys must be unique, and each key is associated with a specific value

	// below we set user1 to be a map having the key as a string type
	// and the values as string types
	user1 := map[string]string {
		"first_name": "John",
		"last_name": "",
	}
	
	// printing the whole map
	fmt.Println(user1)

	// printing just the key
	fmt.Println(user1["first_name"])
	fmt.Println(user1["last_name"])


	// nested map:
	user2 := map[string]map[string]string {
		"name": {
			"first": "Maria",
			"last": "Egies",
		},
		"body_info": {
			"height": "170",
			"weight": "60",
		},
	}

	fmt.Println(user2)

	// adding element
	map1 := map[string]int {
		"one":1,
	}

	map1["two"] = 2
	fmt.Println(map1)

	// deleting element
	delete(map1, "two")
	fmt.Println(map1)
}