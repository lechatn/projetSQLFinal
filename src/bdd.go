package sqlproject


import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)


func OpenDb() *sql.DB { // Function to open the database
	dbPath := "data.db"
	db, errOpenBDD := sql.Open("sqlite3", dbPath)
	if errOpenBDD != nil {
		log.Fatal(errOpenBDD)
	}
	return db
}
