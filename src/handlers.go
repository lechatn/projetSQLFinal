package sqlproject

import (
	"context"
	"database/sql"

	//"fmt"
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
	IdPost        string
	Salary        int
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
	IdProject  string
}

type hierarchy struct {
	IdEmployes string
	IdSuperior string
}

type addEmploye struct {
	DepartementList []departement
	PostList        []post
	Sucess          bool
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

	rows, errQuery2 := db.QueryContext(context.Background(), "SELECT * from employes") // Get the profile picture

	if errQuery2 != nil {
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

	var addemployes addEmploye

	addemployes.DepartementList = departementList
	addemployes.PostList = postList

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

	idPost = idPost[:1]
	idDepartement = idDepartement[:1]

	_, errExec := db.ExecContext(context.Background(), "INSERT INTO employes (name, firstname, birthdate, mail, city, idDepartement, idPost, salary) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", name, firstname, birthdate, mail, city, idDepartement, idPost, salary)

	if errExec != nil {
		http.Error(w, "Error inserting into employes table", http.StatusInternalServerError)
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

func AllProjectsHandler(w http.ResponseWriter, r *http.Request) {}
