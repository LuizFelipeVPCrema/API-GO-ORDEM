package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	migratesqlite "github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	arquivosmigration "github.com/LuizFelipeVPCrema/api-go-ordem/migrations"
)

func ExecutarMigrations(db *sql.DB) error {
	sourceDriver, err := iofs.New(
		arquivosmigration.Arquivos,
		".",
	)
	if err != nil {
		return fmt.Errorf(
			"erro ao carregar arquivos de migration: %w",
			err,
		)
	}

	databaseDriver, err := migratesqlite.WithInstance(
		db,
		&migratesqlite.Config{},
	)
	if err != nil {
		return fmt.Errorf(
			"erro ao configurar driver das migrations: %w",
			err,
		)
	}

	migrator, err := migrate.NewWithInstance(
		"iofs",
		sourceDriver,
		"sqlite",
		databaseDriver,
	)
	if err != nil {
		return fmt.Errorf(
			"erro ao criar executor de migrations: %w",
			err,
		)
	}

	err = migrator.Up()

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf(
			"erro ao executar migrations: %w",
			err,
		)
	}

	return nil
}
