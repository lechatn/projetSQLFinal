package main

import (
	"log"
	"net/http"
	"html/template"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func HomeHandler (w http.ResponseWriter, r *http.Request) {
	tmpl, errReading5 := template.ParseFiles("templates/index.html")
	if errReading5 != nil {
		http.Error(w, "Error reading the HTML file : index.html", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}