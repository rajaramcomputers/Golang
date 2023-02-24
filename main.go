package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB
var Option string

func main() {
	dsn := mysql.Config{
		User:                 "root",
		Passwd:               "mysql",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "sakila",
		AllowNativePasswords: true,
	}
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
	fmt.Println("Connected to MySQL")
	for {
		fmt.Println("Enter the option to perform:\n 1. Add Actor \n 2. Find Actor\n 3. Update Actor \n 4. Show All Records \n 5. Delete Actor:")
		fmt.Scanf("%s", &Option)
		fmt.Printf("Option Selected is %v\n", Option)
		switch Option {
		case "1":
			{
				var fname string
				var lname string
				fmt.Print("Enter the first name:")
				fmt.Scanf("%s", &fname)
				fmt.Print("Enter the last name:")
				fmt.Scanf("%s", &lname)
				fmt.Println(fname, lname)
				actorID, err := addActor(fname, lname)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("ID of added actor %v\n", actorID)
			}
		case "2":
			{
				var tosearch int64
				fmt.Print("Enter the ID to Search:")
				fmt.Scanf("%d", &tosearch)
				actors, err := getActor(int32(tosearch))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Actors Found %v\n", actors)
			}
		case "3":
			{
				var toupdate int32
				var tochangefirstname string
				fmt.Print("Enter the ID to Update:")
				fmt.Scanf("%d", &toupdate)
				fmt.Print("Enter the First Name to Update:")
				fmt.Scanf("%s", &tochangefirstname)

				rowUpdated, err := updateActor(tochangefirstname, toupdate)

				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Total Rows Updated: %d\n", rowUpdated)

			}
		case "4":
			{
				actors, err := getAllActor()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Actors Found %v\n", actors)
			}
		case "5":
			{
				var tosearch int32
				fmt.Print("Enter the ID to Search:")
				fmt.Scanf("%d", &tosearch)
				rowsDeleted, err := deleteActor(tosearch)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Actors Deleted %d\n", rowsDeleted)
			}
		}
		var tocontinue string
		fmt.Print("Enter the y/n to continue:")
		fmt.Scanf("%s", &tocontinue)
		if tocontinue == "n" {
			break
		}

	}
}
func addActor(firstname string, lastname string) (int64, error) {
	result, err := db.Exec("Insert into actor (first_name,last_name) values (?,?)", firstname, lastname)
	if err != nil {
		return 0, fmt.Errorf("addActor: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addActor: %v", err)
	}
	return id, nil
}

type Actor struct {
	actor_id   int64
	first_name string
	last_name  string
}

func getActor(actorID int32) ([]Actor, error) {
	var actors []Actor
	result, err := db.Query("SELECT actor_id, first_name, last_name from actor where actor_id =?", actorID)
	if err != nil {
		return nil, fmt.Errorf("GetActor %v:, %v", actorID, err)
	}
	defer result.Close()
	for result.Next() {
		var acts Actor
		if err := result.Scan(&acts.actor_id, &acts.first_name, &acts.last_name); err != nil {
			return nil, fmt.Errorf("GetActor %v:, %v", actorID, err)
		}
		actors = append(actors, acts)
		if err != nil {
			return nil, fmt.Errorf("GetActor %v: %v", actorID, err)
		}
	}
	return actors, nil
}

func getAllActor() ([]Actor, error) {
	var actors []Actor
	result, err := db.Query("SELECT actor_id, first_name, last_name from actor")
	if err != nil {
		return nil, fmt.Errorf("GetAllActor %v", err)
	}
	defer result.Close()
	for result.Next() {
		var acts Actor
		if err := result.Scan(&acts.actor_id, &acts.first_name, &acts.last_name); err != nil {
			return nil, fmt.Errorf("GetAllActor %v", err)
		}
		actors = append(actors, acts)
		if err != nil {
			return nil, fmt.Errorf("GetAllActor %v", err)
		}
	}
	return actors, nil
}

func updateActor(firstname string, actorid int32) (int64, error) {
	result, err := db.Exec("Update actor set first_name = ? where actor_id = ?", firstname, actorid)
	if err != nil {
		return 0, fmt.Errorf("updateActor : %v", err)
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("updateActor : %v", err)
	}
	return id, nil
}

func deleteActor(actorid int32) (int64, error) {
	result, err := db.Exec("delete from actor where actor_id=?", actorid)
	if err != nil {
		return 0, fmt.Errorf("deleteActor: %v", err)
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("deleteActor: %v", err)
	}
	return id, nil

}
