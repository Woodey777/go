package main

import (
	"database/sql"
	"github.com/lib/pq"
	// "./server"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func main() {
	db, err := sql.Open("postgres", "postgres://localhost:5432/database?sslmode=enable")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
			"file:///migrations",
			"postgres", driver)
	m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
}
