package database

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/tiburciohugo/users-api-go/config/env"
)

func NewDBConnection() (*sql.DB, error) {
	postgresURI := env.Env.DatabaseURL
	db, err := sql.Open("postgres", postgresURI)
	if err != nil {
		slog.Error("Error opening database connection")
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		slog.Error("Error pinging database")
		db.Close()
		return nil, err
	}
	slog.Info("Database connection established", slog.String("package", "database"))
	return db, nil
}
