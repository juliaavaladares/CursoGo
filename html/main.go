package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {
	templates = template.Must(templates.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		user := usuario{
			Nome:  "Julia",
			Email: "julia@gmail.com",
		}

		templates.ExecuteTemplate(w, "home.html", user)
	})

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Carregando pagina de usuarios"))
	})

	fmt.Println("Listen Serve on Port 5000:")
	log.Fatal(
		http.ListenAndServe(":5000", nil),
	)

}
