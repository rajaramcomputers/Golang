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

	if err = modifyActor(db, 30, "JOE"); err != nil {
		log.Fatal(err)
	}
}

func modifyActor(db *sql.DB, actorid int64, firstname string) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	if _, err = db.Exec("UPDATE actor SET first_name = ? WHERE actor_id = ?", firstname, actorid); err != nil {
		return
	}
	return
}
