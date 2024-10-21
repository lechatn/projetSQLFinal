package sqlproject

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"database/sql"
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
		employe.Birthdate = employe.Birthdate[:10]

		employesList = append(employesList, employe)
	}

	errExecute := tmpl.Execute(w, employesList)
	if errExecute != nil {
		log.Printf("Error executing template: %v", errExecute)
		http.Error(w, "Error executing the HTML file : index.html", http.StatusInternalServerError)
		return
	}
}


func AllEmployesHandler (w http.ResponseWriter, r *http.Request) {}

func AddEmployeHandler (w http.ResponseWriter, r *http.Request) {}

func RemoveEmployeHandler (w http.ResponseWriter, r *http.Request) {}

func EditEmployeHandler (w http.ResponseWriter, r *http.Request) {}

func AllProjectsHandler (w http.ResponseWriter, r *http.Request) {}

