package belajar_golang_database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "amrizal:hanyasatu@tcp(localhost:3306)/local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
