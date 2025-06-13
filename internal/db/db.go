package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var pgxConn *pgx.Conn

func GetInstance() (*pgx.Conn, error) {
	if pgxConn != nil {
		return pgxConn, nil
	}

	// This is dummy data
	connStr := "user=postgres dbname=jasonguessr password= sslmode=disable"
	pgxConn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %w", err)
	}

	return pgxConn, nil
}
