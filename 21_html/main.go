package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)

var templates *template.Template

type user struct {
	Name string
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := user{"Bilbo"}
		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}