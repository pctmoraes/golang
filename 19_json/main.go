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
	// function Marshal
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

	// function Unmarshal
	tedJSON := `{"Name":"Ted","Breed":"Husky","Age":6}`

	var ted dog

	if err := json.Unmarshal([]byte(tedJSON), &ted); err != nil {
		log.Fatal(err)
	}
	fmt.Println(ted)


	lunJSON := `{"Name":"Lun","Breed":"Golden","Age":"4"}`

	lun := make(map[string]string)

	// Unmarshal the JSON string into the map
	if errr := json.Unmarshal([]byte(lunJSON), &lun); errr != nil {
		log.Fatal(errr)
	}
	fmt.Println(lun)
}