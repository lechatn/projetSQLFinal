package sqlproject

import (
	"context"
	"fmt"
	"net/http"
)

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
	// Insert into project table the new project
	_, errExec := db.ExecContext(context.Background(), "INSERT INTO project (name, responsable) VALUES (?, ?)", name, responsable)

	if errExec != nil {
		http.Error(w, "Error inserting into project table", http.StatusInternalServerError)
		return
	}
	// Get the idProject of the new project
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
	// Insert into employes_project table the members of the new project
	for _, member := range members {
		_, errExec2 := db.ExecContext(context.Background(), "INSERT INTO employes_project VALUES (?, ?)", member, idProject)

		if errExec2 != nil {
			http.Error(w, "Error inserting into employes_project table", http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/allprojects", http.StatusSeeOther)
}