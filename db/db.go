package db

import (
	"Invoices/models"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUrl  string = "proxymanger:YaABs8NW@tcp(localhost:3306)/proxymanger"
	server string = "mysql"
)

var sessionDB *sql.DB
var err error

func init() {
	sessDB,err := sql.Open(server,dbUrl)
	if err != nil {
		panic(err)
	}
	sessionDB = sessDB
	defer sessionDB.Close()
}

func UpdateDepartment(d models.Department, m string) error {
	switch m {
	case "delete":
		//delete from db
	case "update":
		// update record
	case "add":
		//add record
	default:
		err = errors.New("Command Not Found")
	}
	return err
}