package main

import (
	"fmt"
	"encoding/json"
	"log"
	"bytes"
)

type dog struct {
	Name string
	Breed string
	Age uint
}

func main() {
	meg := dog{"Meg", "Lhasa apso", 7}
	fmt.Println(meg)

	megJSON, _error := json.Marshal(meg)
	if _error != nil {
		log.Fatal(_error)
	}

	fmt.Println(bytes.NewBuffer(megJSON))

	nina := map[string]string{
		"Name": "Nina", 
		"Breed": "Yorkshire",
		"Age": "Three",
	}

	fmt.Println(nina)

	ninaJSON, _error := json.Marshal(nina)
	if _error != nil {
		log.Fatal(_error)
	}

	fmt.Println(bytes.NewBuffer(ninaJSON))

}