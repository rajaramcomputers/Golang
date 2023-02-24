package main

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// a successful case
func Test_modifyActor_PASS(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening database connection", err)
	}
	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE actor").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	// now we execute our method
	if err = modifyActor(db, 50, "JOE"); err != nil {
		t.Errorf("Error was not expected while modifing actors: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func Test_modifyActor_FAIL(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening database connection", err)
	}
	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE actor").WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()
	// now we execute our method
	if err = modifyActor(db, 50, "JOE"); err == nil {
		t.Errorf("Error was expected while modifing actors but no error occured")
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
