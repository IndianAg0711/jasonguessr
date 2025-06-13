package db

import (
	"github.com/golang-migrate/migrate/v4"
)

func Migrate() {
	dbConn, _ := GetInstance()

	// need to match up the migrate db type with our pgx connection (check https://github.com/golang-migrate/migrate/tree/v4.18.3/database/pgx/v5)
	m, err := migrate.NewWithDatabaseInstance(
		"file",
		"jasonguessr",
		dbConn,
	)
}
