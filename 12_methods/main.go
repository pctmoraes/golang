package main

import "fmt"

type user struct {
	name  string
	age uint8
}

// methods are functions associated with a specific type
// unlike some other object-oriented programming languages, 
// Go does not have classes, so instead, it has types
// and methods can be associated with these types.

func (u user) save() {
	fmt.Printf("Saving user %s data on db\n", u.name)
}

func (u user) isUserOfLegalAge() bool {
	return u.age >= 18
}

func (u *user) hasBirthday() {
	u.age++
}

func main() {
	user1 := user{"Mila", 20}
	user1.save()

	user2 := user{"Davi", 30}
	user2.save()

	isUserOfLegalAge := user2.isUserOfLegalAge()
	fmt.Println(isUserOfLegalAge)

	user2.hasBirthday()
	fmt.Println(user2.age)

}