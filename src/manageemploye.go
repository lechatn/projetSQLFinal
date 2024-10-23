package sqlproject

import (
	"context"
	"html/template"
	"log"
	"net/http"
	
)

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
	tmpl, errReading4 := template.ParseFiles("templat/edit.html")
	if errReading4 != nil {
		http.Error(w, "Error reading the HTML file : allEmployes.html", http.StatusInternalServerError)
		return
	}

	errExecute := tmpl.Execute(w, nil)
	if errExecute != nil {
		log.Printf("Error executing template: %v", errExecute)
		http.Error(w, "Error executing the HTML file : manage.html", http.StatusInternalServerError)
		return
	}

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

	_, errExec := db.ExecContext(context.Background(), "DELETE FROM employes WHERE idEmployes = ?", IdEmployes)
	
	if errExec != nil {
		http.Error(w, "Error deleting employe", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}


func EditEmployeHandler(w http.ResponseWriter, r *http.Request) {}


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
