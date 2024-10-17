package main

import (
	"context"
	"database/sql"
	//"fmt"
	"html/template"
	"log"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type employes struct {
	IdEmployes string
	Name string
	Firstname string
	Birthdate string
	Mail string
	City string
	IdDepartement string
	IdPost string
	Salary int
}

type departement struct {
	IdDepartement string
	Name string
}

type post struct {
	IdPost string
	Name string
}

type project struct {
	IdProject string
	Name string
	Responsable string
}

type employes_project struct {
	IdEmployes string
	IdProject string
}

type hierarchy struct {
	IdEmployes string
	IdSuperior string
}


func main() {
	http.HandleFunc("/", HomeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func HomeHandler (w http.ResponseWriter, r *http.Request) {
	db = OpenDb()
	tmpl, errReading1 := template.ParseFiles("templates/index.html")
	if errReading1 != nil {
		http.Error(w, "Error reading the HTML file : index.html", http.StatusInternalServerError)
		return
	}

	rows, errQuery19 := db.QueryContext(context.Background(), "SELECT * from employes") // Get the profile picture

	if errQuery19 != nil {
		http.Error(w, "Error with employes table", http.StatusInternalServerError)
		return
	}

	if rows != nil {
		defer rows.Close()
	}

	var employesList []employes

	for rows.Next() {
		var employe employes
		errScan := rows.Scan(&employe.IdEmployes, &employe.Name, &employe.Firstname, &employe.Birthdate, &employe.Mail, &employe.City, &employe.IdDepartement, &employe.IdPost, &employe.Salary)
		if errScan != nil {
			http.Error(w, "Error with employes table", http.StatusInternalServerError)
			return
		}

		employesList = append(employesList, employe)
	}

	errExecute := tmpl.Execute(w, employesList)
	if errExecute != nil {
		log.Printf("Error executing template: %v", errExecute)
		http.Error(w, "Error executing the HTML file : index.html", http.StatusInternalServerError)
		return
	}
}

func OpenDb() *sql.DB { // Function to open the database
	dbPath := "data.db"
	db, errOpenBDD := sql.Open("sqlite3", dbPath)
	if errOpenBDD != nil {
		log.Fatal(errOpenBDD)
	}
	return db
}

