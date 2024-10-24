package sqlproject

import (
	"context"
	"html/template"
	"log"
	"net/http"
)

func AllProjectsHandler (w http.ResponseWriter, r *http.Request) {
	db = OpenDb()

	tmpl, errReading2 := template.ParseFiles("templates/allProjects.html")
	if errReading2 != nil {
		http.Error(w, "Error reading the HTML file : allProjects.html", http.StatusInternalServerError)
		return
	}
    //Get all data of projects and their responsable name
	rows, errQuery2 := db.QueryContext(context.Background(),
		 `SELECT project.idProject, project.name, employes.name || ' ' || employes.firstname AS responsable
		 FROM project
		 JOIN employes ON project.responsable = employes.idEmployes
		 `)

	if errQuery2 != nil {
		http.Error(w, "Error with project table", http.StatusInternalServerError)
		log.Printf("Error with project table in query: %v", errQuery2)
		return
	}

	if rows != nil {
		defer rows.Close()
	}

	var projectList []project

	for rows.Next() {
		var project project
		errScan := rows.Scan(&project.IdProject, &project.Name, &project.Responsable)
		if errScan != nil {
			http.Error(w, "Error with project table", http.StatusInternalServerError)
			log.Printf("Error scanning project table: %v", errScan)
			return
		}

		projectList = append(projectList, project)

	}

	//Get all data of employes_project and their member name 
	rows2, errQuery3 := db.QueryContext(context.Background(),
	 `SELECT employes_project.*, employes.name || ' ' || employes.firstname AS member
	 FROM employes_project
	 JOIN employes ON employes_project.idEmployes = employes.idEmployes
	 `)

	if errQuery3 != nil {
		http.Error(w, "Error with employes_project table", http.StatusInternalServerError)
		log.Printf("Error with employes_project table in query: %v", errQuery3)
		return
	}

	if rows2 != nil {
		defer rows2.Close()
	}

	var membersList []employes_project

	for rows2.Next() {
		var member employes_project
		errScan := rows2.Scan(&member.IdEmployes, &member.IdProject, &member.MemberName)
		if errScan != nil {
			http.Error(w, "Error with employes_project table", http.StatusInternalServerError)
			log.Printf("Error scanning employes_project table: %v", errScan)
			return

		}

		membersList = append(membersList, member)

	}
	//Get some data of employes
	rows3, errQuery4 := db.QueryContext(context.Background(), `SELECT name, firstname, idEmployes FROM employes`)

	if errQuery4 != nil {
		http.Error(w, "Error with employes table in query", http.StatusInternalServerError)
		return
	}

	if rows3 != nil {
		defer rows3.Close()
	}

	var employesList []employes

	for rows3.Next() {
		var employe employes
		errScan := rows3.Scan(&employe.Name, &employe.Firstname, &employe.IdEmployes)
		if errScan != nil {
			http.Error(w, "Error with employes table", http.StatusInternalServerError)
			return
		}

		employesList = append(employesList, employe)

	}


	var allprojects allProjects

	allprojects.ProjectList = projectList
	allprojects.Members = membersList
	allprojects.Employes = employesList

	errExecute := tmpl.Execute(w,allprojects)
	if errExecute != nil {
		log.Printf("Error executing template: %v", errExecute)
		http.Error(w, "Error executing the HTML file : index.html", http.StatusInternalServerError)
		return
	}
}