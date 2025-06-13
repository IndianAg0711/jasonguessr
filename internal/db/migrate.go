package db

import (
	// "pgx"
	"github.com/IndianAg0711/jasonguessr/internal/db"
	"github.com/jackc/pgx/v5"
)

func Migrate(){
	driver, err := pgx.WithInstance(db.GetInstance())
}
