package sqlproject

import (
	"database/sql" //"fmt"
	"html/template"
	"log"
	"net/http"
)

var db *sql.DB

type employes struct {
	IdEmployes    string
	Name          string
	Firstname     string
	Birthdate     string
	Mail          string
	City          string
	IdDepartement string
	IdPost string
	Salary int
	DepartementName string
	PostName string
}

type departement struct {
	IdDepartement string
	Name          string
}

type post struct {
	IdPost string
	Name   string
}

type project struct {
	IdProject   string
	Name        string
	Responsable string
}

type employes_project struct {
	IdEmployes string
	IdProject string
	MemberName string
}

type hierarchy struct {
	IdEmployes string
	IdSuperior string
}

type addEmploye struct {
	DepartementList []departement
	PostList []post
	EmployeList []employes
	ProjectList []project
}

type allProjects struct {
	ProjectList []project
	Members []employes_project
	Employes []employes
}

type editEmploye struct {
	DepartementList []departement
	PostList        []post
	Employe         employes
}


func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, errReading1 := template.ParseFiles("templates/index.html")
	if errReading1 != nil {
		http.Error(w, "Error reading the HTML file : index.html", http.StatusInternalServerError)
		return
	}

	errExecute := tmpl.Execute(w, nil)
	if errExecute != nil {
		log.Printf("Error executing template: %v", errExecute)
		http.Error(w, "Error executing the HTML file : index.html", http.StatusInternalServerError)
		return
	}
}

