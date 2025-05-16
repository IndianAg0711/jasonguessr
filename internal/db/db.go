package db

import (
	"database/sql"
)

var dbInstance *sql.DB

func GetInstance() *sql.DB {
	if dbInstance != nil {
		return dbInstance
	}

	// This is dummy data
	connStr := "user=myuser dbname=mydb password=mypassword sslmode=disable"
	dbInstance, _ := sql.Open("postgres", connStr)

	return dbInstance
}
