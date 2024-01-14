package main

import "fmt"

type user struct {
	name string
	email string
	age uint
}

func main () {
	// initializes the user1 var and go will set 
	// the user fields with default values
	var user1 user
	fmt.Println(user1)

	// set values to the user1 var
	user1.name = "Paula"
	user1.email = "paula@email.com"
	user1.age = 29
	fmt.Println(user1)


	// initializes and set values to the user2 in all fields
	user2 := user{"Joao", "joao@email.com", 15}
	fmt.Println(user2)


	// initializes and set value to the user2 in one field
	user3 := user{name: "Amber"}
	fmt.Println(user3)
}