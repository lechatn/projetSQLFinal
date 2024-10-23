package sqlproject

import (
	"context"
	"net/http"
	"html/template"
	"log"
)

func SubmitEmployeHandler(w http.ResponseWriter, r *http.Request) {
	db = OpenDb()
	// Get the form values
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	errParse := r.ParseForm()

	if errParse != nil {
		http.Error(w, "Error parsing the form", http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")
	firstname := r.FormValue("firstname")
	birthdate := r.FormValue("birthdate")
	mail := r.FormValue("mail")
	city := r.FormValue("city")
	idDepartement := r.FormValue("departement")
	idPost := r.FormValue("post")
	salary := r.FormValue("salary")
	superior := r.FormValue("superior")
	project := r.FormValue("project")

	idPost = idPost[:1]
	idDepartement = idDepartement[:1]

	//Insert new employe into employes table
	_, errExec := db.ExecContext(context.Background(), "INSERT INTO employes (name, firstname, birthdate, mail, city, idDepartement, idPost, salary) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", name, firstname, birthdate, mail, city, idDepartement, idPost, salary)

	if errExec != nil {
		http.Error(w, "Error inserting into employes table", http.StatusInternalServerError)
		return
	}
	//Get the idEmployes of the new employe
	rows6, errQuery6 := db.QueryContext(context.Background(), `SELECT idEmployes FROM employes WHERE name = ? AND firstname = ?`, name, firstname)

	if errQuery6 != nil {
		http.Error(w, "Error with employes table in query", http.StatusInternalServerError)
		return
	}

	if rows6 != nil {
		defer rows6.Close()
	}

	var idEmployes string

	for rows6.Next() {
		errScan := rows6.Scan(&idEmployes)
		if errScan != nil {
			http.Error(w, "Error with employes table", http.StatusInternalServerError)
			return
		}

	}
	//Insert the new employe into hierarchy table with his superior
	_, errExec2 := db.ExecContext(context.Background(), "INSERT INTO hierarchy (idEmployes, idSuperior) VALUES (?, ?)", idEmployes, superior)

	if errExec2 != nil {
		http.Error(w, "Error inserting into hierarchy table", http.StatusInternalServerError)
		return
	}
	//Insert the new employe into employes_project table with his project
	_, errExec3 := db.ExecContext(context.Background(), "INSERT INTO employes_project VALUES (?, ?)", idEmployes, project)

	if errExec3 != nil {
		http.Error(w, "Error inserting into employes_project table", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/addemploye", http.StatusSeeOther)
}


func AddEmployeHandler(w http.ResponseWriter, r *http.Request) {
	db = OpenDb()
	tmpl, errReading3 := template.ParseFiles("templates/addEmploye.html")
	if errReading3 != nil {
		http.Error(w, "Error reading the HTML file : addEmploye.html", http.StatusInternalServerError)
		return
	}
	//Get all the data from the table departement
	rows4, errQuery4 := db.QueryContext(context.Background(), "SELECT * from departement")

	if errQuery4 != nil {
		http.Error(w, "Error with departement table in query", http.StatusInternalServerError)
		return
	}

	if rows4 != nil {
		defer rows4.Close()
	}

	var departementList []departement

	for rows4.Next() {
		var depart departement
		errScan := rows4.Scan(&depart.IdDepartement, &depart.Name)
		if errScan != nil {
			http.Error(w, "Error with departement table", http.StatusInternalServerError)
			return
		}

		departementList = append(departementList, depart)
	}
	//Get all the data from the table post
	rows5, errQuery5 := db.QueryContext(context.Background(), "SELECT * from post")

	if errQuery5 != nil {
		http.Error(w, "Error with post table in query", http.StatusInternalServerError)
		return
	}

	if rows5 != nil {
		defer rows5.Close()
	}

	var postList []post

	for rows5.Next() {
		var post post
		errScan := rows5.Scan(&post.IdPost, &post.Name)
		if errScan != nil {
			http.Error(w, "Error with post table", http.StatusInternalServerError)
			return
		}

		postList = append(postList, post)

	}
	//Get some data from the table employes
	rows6, errQuery6 := db.QueryContext(context.Background(), `SELECT idEmployes, name, firstname FROM employes`)

	if errQuery6 != nil {
		http.Error(w, "Error with employes table in query", http.StatusInternalServerError)
		return
	}

	if rows6 != nil {
		defer rows6.Close()
	}

	var employeList []employes

	for rows6.Next() {
		var employe employes
		errScan := rows6.Scan(&employe.IdEmployes, &employe.Name, &employe.Firstname)
		if errScan != nil {
			http.Error(w, "Error with employes table", http.StatusInternalServerError)
			return
		}

		employeList = append(employeList, employe)

	}

	//Get some data from the table project
	rows7, errQuery7 := db.QueryContext(context.Background(), `SELECT idProject, name FROM project`)

	if errQuery7 != nil {
		http.Error(w, "Error with project table in query", http.StatusInternalServerError)
		return
	}

	if rows7 != nil {
		defer rows7.Close()
	}

	var projectList []project

	for rows7.Next() {
		var project project
		errScan := rows7.Scan(&project.IdProject, &project.Name)
		if errScan != nil {
			http.Error(w, "Error with project table", http.StatusInternalServerError)
			return
		}

		projectList = append(projectList, project)

	}

	var addemployes addEmploye

	addemployes.DepartementList = departementList
	addemployes.PostList = postList
	addemployes.EmployeList = employeList
	addemployes.ProjectList = projectList



	errExecute := tmpl.Execute(w, addemployes)
	if errExecute != nil {
		log.Printf("Error executing template: %v", errExecute)
		http.Error(w, "Error executing the HTML file : addEmploye.html", http.StatusInternalServerError)
		return
	}
}