package main

import "fmt"

// In this exercise we will understand about "heritage" in go
// it's not like in OOP languages but is the way go has implemented it
// the only diff between this exercise and the exercise on the script
// user_with_address.go is that when passing the address struct to
// the user struct we have to call by the struct's name \/

type user struct {
	name string
	email string
	age uint
	// like this
	address
}

type address struct {
	street string
	number uint
	neighbourhood string
	city string
	district string
}

// and we can access the address fields directly without having
// to pass the "address" name \/

func (user1 user) String() string {
	return fmt.Sprintf(
		"Name: %s\nEmail: %s\nAge: %d\nAddress: %s, %d, %s, %s - %s", 
		user1.name, user1.email, user1.age, 
		// like this
		user1.street, user1.number, user1.neighbourhood, 
		user1.city, user1.district,
	)
}

func main () {
	address1 := address {
		"Dummy Street",
		90,
		"Berghud",
		"Z City",
		"EX",
	}

	user1 := user{
		"Amber", 
		"amber@email.com", 
		22,
		address1,
	}

	fmt.Println(user1.String())
}