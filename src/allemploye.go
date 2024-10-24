package sqlproject

import (
	"context"
	"html/template"
	"log"
	"net/http"
)

func AllEmployesHandler(w http.ResponseWriter, r *http.Request) {
	db = OpenDb()

	tmpl, errReading2 := template.ParseFiles("templates/allEmployes.html")
	if errReading2 != nil {
		http.Error(w, "Error reading the HTML file : allEmployes.html", http.StatusInternalServerError)
		return
	}
	//Get all employes with their departement name and post name
	rows, errQuery2 := db.QueryContext(context.Background(), `
        SELECT employes.*, departement.name, post.name
		FROM employes
		JOIN departement ON employes.idDepartement = departement.idDepartement
		JOIN post ON employes.idPost = post.idPost
    `)

	if errQuery2 != nil {
		http.Error(w, "Error with employes table", http.StatusInternalServerError)
		log.Printf("Error with employes table in query: %v", errQuery2)
		return
	}

	if rows != nil {
		defer rows.Close()
	}

	var employesList []employes

	for rows.Next() {
		var employe employes
		errScan := rows.Scan(&employe.IdEmployes, &employe.Name, &employe.Firstname, &employe.Birthdate, &employe.Mail, &employe.City, &employe.IdDepartement, &employe.IdPost, &employe.Salary, &employe.DepartementName, &employe.PostName)
		if errScan != nil {
			http.Error(w, "Error with employes table", http.StatusInternalServerError)
			log.Printf("Error scanning employes table: %v", errScan)
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
