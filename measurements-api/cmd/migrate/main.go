package main

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"measurements-api/internal/config"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatal(err, "failed to load config")
	}

	cfg := config.FromEnv()
	m, err := migrate.New("file://./cmd/migrate/migrations/postgresql", cfg.Postgres.URL())
	if err != nil {
		log.Fatal(err)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	version, dirty, err := m.Version()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Applied migration: %d. Dirty: %t\n", version, dirty)
}
