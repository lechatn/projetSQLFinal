package sqlproject

import (
	"context"
	"database/sql" //"fmt"
	"fmt"
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

func ManageHandler(w http.ResponseWriter, r *http.Request) {
	db = OpenDb()

	tmpl, errReading2 := template.ParseFiles("templates/manage.html")
	if errReading2 != nil {
		http.Error(w, "Error reading the HTML file : manage.html", http.StatusInternalServerError)
		return
	}

	rows, errQuery2 := db.QueryContext(context.Background(), "SELECT idEmployes,name,firstname FROM employes")

	if errQuery2 != nil {
		http.Error(w, "Error with employes table 1", http.StatusInternalServerError)
		return
	}

	if rows != nil {
		defer rows.Close()
	}

	var employesList []employes

	for rows.Next() {
		var employe employes
		errScan := rows.Scan(&employe.IdEmployes, &employe.Name, &employe.Firstname)
		if errScan != nil {
			http.Error(w, "Error with employes table 2", http.StatusInternalServerError)
			return
		}

		employesList = append(employesList, employe)
	}

	errExecute := tmpl.Execute(w, employesList)
	if errExecute != nil {
		log.Printf("Error executing template: %v", errExecute)
		http.Error(w, "Error executing the HTML file : manage.html", http.StatusInternalServerError)
		return
	}
}

func AllEmployesHandler(w http.ResponseWriter, r *http.Request) {
	db = OpenDb()

	tmpl, errReading2 := template.ParseFiles("templates/allEmployes.html")
	if errReading2 != nil {
		http.Error(w, "Error reading the HTML file : allEmployes.html", http.StatusInternalServerError)
		return
	}

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

func AddEmployeHandler(w http.ResponseWriter, r *http.Request) {
	db = OpenDb()
	tmpl, errReading3 := template.ParseFiles("templates/addEmploye.html")
	if errReading3 != nil {
		http.Error(w, "Error reading the HTML file : addEmploye.html", http.StatusInternalServerError)
		return
	}

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

	_, errExec := db.ExecContext(context.Background(), "INSERT INTO employes (name, firstname, birthdate, mail, city, idDepartement, idPost, salary) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", name, firstname, birthdate, mail, city, idDepartement, idPost, salary)

	if errExec != nil {
		http.Error(w, "Error inserting into employes table", http.StatusInternalServerError)
		return
	}

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

	_, errExec2 := db.ExecContext(context.Background(), "INSERT INTO hierarchy (idEmployes, idSuperior) VALUES (?, ?)", idEmployes, superior)

	if errExec2 != nil {
		http.Error(w, "Error inserting into hierarchy table", http.StatusInternalServerError)
		return
	}

	_, errExec3 := db.ExecContext(context.Background(), "INSERT INTO employes_project VALUES (?, ?)", idEmployes, project)

	if errExec3 != nil {
		http.Error(w, "Error inserting into employes_project table", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/addemploye", http.StatusSeeOther)
}

func RemoveHandler(w http.ResponseWriter, r *http.Request) {
	db = OpenDb()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	errParse := r.ParseForm()

	if errParse != nil {
		http.Error(w, "Error parsing the form", http.StatusInternalServerError)
		return
	}

	IdEmployes := r.FormValue("idremove")

	_, errExec := db.ExecContext(context.Background(), "DELETE FROM employes WHERE idEmployes = ?", IdEmployes)
	
	if errExec != nil {
		http.Error(w, "Error deleting employe", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/manage", http.StatusSeeOther)

}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	db = OpenDb()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	errParse := r.ParseForm()

	if errParse != nil {
		http.Error(w, "Error parsing the form", http.StatusInternalServerError)
		return
	}

	IdEmployes := r.FormValue("idedit")

	rows, errQuery := db.QueryContext(context.Background(), "SELECT * from employes WHERE idEmployes = ?", IdEmployes)

	if errQuery != nil {
		http.Error(w, "Error with employes table", http.StatusInternalServerError)
		return
	}

	if rows != nil {
		defer rows.Close()
	}

	var employe employes

	for rows.Next() {
		errScan := rows.Scan(&employe.IdEmployes, &employe.Name, &employe.Firstname, &employe.Birthdate, &employe.Mail, &employe.City, &employe.IdDepartement, &employe.IdPost, &employe.Salary)
		if errScan != nil {
			http.Error(w, "Error with employes table", http.StatusInternalServerError)
			return
		}

		employe.Birthdate = employe.Birthdate[:10]
	}

	tmpl, errReading4 := template.ParseFiles("templates/editEmploye.html")

	if errReading4 != nil {
		http.Error(w, "Error reading the HTML file : editEmploye.html", http.StatusInternalServerError)
		return
	}

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

	var editemployes editEmploye

	editemployes.DepartementList = departementList
	editemployes.PostList = postList
	editemployes.Employe = employe

	errExecute := tmpl.Execute(w, editemployes)
	if errExecute != nil {
		log.Printf("Error executing template: %v", errExecute)
		http.Error(w, "Error executing the HTML file : editEmploye.html", http.StatusInternalServerError)
		return
	}

}
func EditEmployeHandler(w http.ResponseWriter, r *http.Request) {
	db = OpenDb()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	errParse := r.ParseForm()

	if errParse != nil {
		http.Error(w, "Error parsing the form", http.StatusInternalServerError)
		return
	}

	IdEmployes := r.FormValue("idemploye")
	name := r.FormValue("name")
	firstname := r.FormValue("firstname")
	birthdate := r.FormValue("birthdate")
	mail := r.FormValue("mail")
	city := r.FormValue("city")
	idDepartement := r.FormValue("departement")
	idPost := r.FormValue("post")
	salary := r.FormValue("salary")

	idPost = idPost[:1]
	idDepartement = idDepartement[:1]

	_, errExec := db.ExecContext(context.Background(), "UPDATE employes SET name = ?, firstname = ?, birthdate = ?, mail = ?, city = ?, idDepartement = ?, idPost = ?, salary = ? WHERE idEmployes = ?", name, firstname, birthdate, mail, city, idDepartement, idPost, salary, IdEmployes)

	if errExec != nil {
		http.Error(w, "Error updating employe", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/manage", http.StatusSeeOther)
}

func AllProjectsHandler (w http.ResponseWriter, r *http.Request) {
	db = OpenDb()

	tmpl, errReading2 := template.ParseFiles("templates/allProjects.html")
	if errReading2 != nil {
		http.Error(w, "Error reading the HTML file : allProjects.html", http.StatusInternalServerError)
		return
	}

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

func AddProjectHandler (w http.ResponseWriter, r *http.Request) {
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
	responsable := r.FormValue("responsable")
	members := r.Form["employes[]"]

	fmt.Println(responsable)
	fmt.Println(members)

	_, errExec := db.ExecContext(context.Background(), "INSERT INTO project (name, responsable) VALUES (?, ?)", name, responsable)

	if errExec != nil {
		http.Error(w, "Error inserting into project table", http.StatusInternalServerError)
		return
	}

	rows6, errQuery6 := db.QueryContext(context.Background(), `SELECT idProject FROM project WHERE name = ?`, name)

	if errQuery6 != nil {
		http.Error(w, "Error with project table in query", http.StatusInternalServerError)
		return
	}

	if rows6 != nil {
		defer rows6.Close()
	}

	var idProject string

	for rows6.Next() {
		errScan := rows6.Scan(&idProject)
		if errScan != nil {
			http.Error(w, "Error with project table", http.StatusInternalServerError)
			return
		}

	}

	for _, member := range members {
		_, errExec2 := db.ExecContext(context.Background(), "INSERT INTO employes_project VALUES (?, ?)", member, idProject)

		if errExec2 != nil {
			http.Error(w, "Error inserting into employes_project table", http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/allprojects", http.StatusSeeOther)
}