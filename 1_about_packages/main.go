package main

import (
	"fmt"
	"about_packages/my_packages"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("I am main print")
	a_package.FirstPackagePrint()
	a_package.SecondPackagePrint()

	// Below is the second exercise about packages
	// where i am going to use a third-party library
	// here is the library link: https://github.com/badoux/checkmail
	// it's a Golang package for email validation
	checkmail_error := checkmail.ValidateFormat("my_email@email.com")
	
	if checkmail_error != nil {
		fmt.Println(checkmail_error)
	}
		
}
