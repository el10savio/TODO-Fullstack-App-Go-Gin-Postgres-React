package api

import (
	"database/sql"
	"fmt"
	"testing"
)

var db_test *sql.DB
var err_test error

func TestMain(m *testing.M) {

	db_test, err_test = sql.Open("postgres", "postgres://postgres:password@localhost/todo?sslmode=disable")

	if err_test != nil {
		fmt.Println(err_test.Error())
		panic(err_test)
	}

	if err_test = db_test.Ping(); err_test != nil {
		fmt.Println(err_test.Error())
		panic(err_test)
	}

}
