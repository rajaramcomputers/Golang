package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

var db *sql.DB

func main() {
	// Data Source Name Properties
	ConnString := "server=localhost;user id=sa;password=sql;port=1433;database=AdventureWorks;"

	// Get a database handle
	var err error
	db, err = sql.Open("sqlserver", ConnString)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
