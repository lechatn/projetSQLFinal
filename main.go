package main

import (
	//"fmt"
	"log"
	"net/http"
	sqlproject "sqlproject/src"
	_ "github.com/mattn/go-sqlite3"
)


func main() {
	http.HandleFunc("/", sqlproject.HomeHandler)
	http.HandleFunc("/allemployes", sqlproject.AllEmployesHandler)
	http.HandleFunc("/addemploye", sqlproject.AddEmployeHandler)
	http.HandleFunc("/manage", sqlproject.ManageHandler)
	http.HandleFunc("/editemploye", sqlproject.EditEmployeHandler)
	http.HandleFunc("/allprojects", sqlproject.AllProjectsHandler)
	http.HandleFunc("/submitemploye", sqlproject.SubmitEmployeHandler)
	http.HandleFunc("/addproject", sqlproject.AddProjectHandler)
	http.HandleFunc("/remove", sqlproject.RemoveHandler)
	http.HandleFunc("/edit", sqlproject.EditHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

