package main

import (
	"context"
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

	ctx := context.Background()
	actorID, err := txActor(ctx, "JOEN", "BERRY")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Actor ID: %v\n", actorID)
}

func txActor(ctx context.Context, firstname string, lastname string) (int64, error) {

	// Begin Transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("Adding Actor Failed: %v", err)
	}
	// Defer a rollback in case of failure
	defer tx.Rollback()

	// Check if name exists
	var actID int64
	if err = tx.QueryRowContext(ctx, "SELECT actor_id from actor where first_name = ? and last_name = ?",
		firstname, lastname).Scan(&actID); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Actor does not exist")
		} else {
			return 0, fmt.Errorf("txActor: %v", err)
		}
	}
	// Rollback if actor exists
	if actID > 0 {
		if err = tx.Rollback(); err != nil {
			return 0, fmt.Errorf("txActor: %v", err)
		}
		fmt.Println("Actor already exist: ", actID)
		fmt.Println("*** Transaction Rolling Back ***")
		return actID, nil
	}

	// Create a new row
	result, err := tx.ExecContext(ctx, "INSERT INTO actor (first_name, last_name) VALUES (?, ?)",
		firstname, lastname)
	if err != nil {
		return 0, fmt.Errorf("txActor: %v", err)
	}

	// Get the ID of the actor just inserted
	NewActorID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("txActor: %v", err)
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return 0, fmt.Errorf("txActor: %v", err)
	} else {
		fmt.Println("New Actor Created: ", NewActorID)
		fmt.Println("*** Transaction Commited ***")
	}

	return NewActorID, nil
}
