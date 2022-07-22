package adapters

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
)

//NewDatabase return db connection and ensure db schema
func NewDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./wallet.db")
	if err != nil {
		return nil, err
	}
	if err := ensureSchema(db); err != nil {
		log.Fatalln("migration failed ", err)
	}

	return db, nil
}

//go:embed migrations
var migrations embed.FS

//ensureSchema creates database schema needed to run the application
func ensureSchema(db *sql.DB) error {
	sourceInstance, err := httpfs.New(http.FS(migrations), "migrations")
	if err != nil {
		return fmt.Errorf("invalid source instance, %w", err)
	}
	targetInstance, err := sqlite3.WithInstance(db, new(sqlite3.Config))
	if err != nil {
		return fmt.Errorf("invalid target sqlite instance, %w", err)
	}
	m, err := migrate.NewWithInstance(
		"httpfs", sourceInstance, "sqlite3", targetInstance)
	if err != nil {
		return fmt.Errorf("failed to initialize migrate instance, %w", err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return sourceInstance.Close()
}
