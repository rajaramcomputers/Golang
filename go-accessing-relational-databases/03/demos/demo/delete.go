package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Data Source Name Properties
	dsn := mysql.Config{
		User:                 "root",
		Passwd:               "MySQL",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "sakila",
		AllowNativePasswords: true,
	}

	// Get a database handle
	var err error
	db, err = sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	rowsDeleted, err := deleteActor(201)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total actor deleted: %d\n", rowsDeleted)
}

func deleteActor(actorid int32) (int64, error) {
	result, err := db.Exec("DELETE FROM actor WHERE actor_id = ?",
		actorid)
	if err != nil {
		return 0, fmt.Errorf("deleteActor: %v", err)
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("deleteActor: %v", err)
	}
	return id, nil
}
