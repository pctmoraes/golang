package main

import (
	"fmt"
	"time"
)

func main() {
	// loops are repeating structures and in golang there is only for loops

	// bellow are the many ways you can use for loops

	// 1st way
	i := 0

	for i < 5 {
		i++
		fmt.Println("Incrementing 1 in i")
		fmt.Println("i:", i)
		time.Sleep(time.Second)
	}


	// 2nd way
	for j := 0; j < 5; j++ {
		fmt.Println("Incrementing 1 in j")
		fmt.Println("j:", j)
		time.Sleep(time.Second)
	}

	// 3rd way
	names := []string{"Lucy","Anakin","Kent"}

	for _, name := range names { // the underscore is the index
		fmt.Println(name)
		time.Sleep(time.Second)
	}

	// 4th way
	users := map[string]string {
		"lucas@email.com": "GHziY687",
		"mia@email.com": "G098OIiY687",
	}

	for key, value := range users {
		fmt.Println(key,value)
		time.Sleep(time.Second)
	}
}