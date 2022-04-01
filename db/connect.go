package db

import (
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx/v5/pgxpool"
)

var connectionPool * pgxpool.Pool

func Connect() {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("BASIS_DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	connectionPool = pool
}