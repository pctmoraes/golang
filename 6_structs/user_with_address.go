package main

import "fmt"

type user struct {
	name string
	email string
	age uint
	address_ address
}

type address struct {
	street string
	number uint
	neighbourhood string
	city string
	district string
}

func (user1 user) String() string {
	return fmt.Sprintf(
		"Name: %s\nEmail: %s\nAge: %d\nAddress: %s, %d, %s, %s - %s", 
		user1.name, user1.email, user1.age, user1.address_.street, 
		user1.address_.number, user1.address_.neighbourhood, 
		user1.address_.city, user1.address_.district,
	)
}

func main () {
	address1 := address {
		"Silly Avenue",
		10,
		"Berghod",
		"Z City",
		"EX",
	}

	user1 := user{
		"John", 
		"john@email.com", 
		15,
		address1,
	}

	fmt.Println(user1.String())
}