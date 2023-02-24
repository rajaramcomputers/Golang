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

	actors, err := GetActorSP("RIVER", "DEAN")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Actor found: %v\n", actors)
}

type Actor struct {
	actor_id   int64
	first_name string
	last_name  string
}

func GetActorSP(first_name string, last_name string) ([]Actor, error) {
	var actors []Actor

	result, err := db.Query("CALL addActor (?, ?)", first_name, last_name)
	if err != nil {
		return nil, fmt.Errorf("GetActorSP: %v", err)
	}
	defer result.Close()

	// Loop through rows
	for result.Next() {
		var acts Actor
		if err := result.Scan(&acts.actor_id, &acts.first_name, &acts.last_name); err != nil {
			return nil, fmt.Errorf("GetActorSP: %v", err)
		}
		actors = append(actors, acts)

		if err := result.Err(); err != nil {
			return nil, fmt.Errorf("GetActorSP: %v", err)
		}
	}
	return actors, nil
}
