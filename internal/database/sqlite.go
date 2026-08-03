package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	_ "modernc.org/sqlite"
)

func ConectarSQLite(path string) (*sql.DB, error) {
	diretorioBanco := filepath.Dir(path)

	if err := os.MkdirAll(diretorioBanco, 0755); err != nil {
		return nil, fmt.Errorf(
			"erro ao criar diretorio do banco: %w",
			err,
		)
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf(
			"erro ao abrir banco SQLite: %w",
			err,
		)
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()

		return nil, fmt.Errorf(
			"erro ao validar conexao com SQLite: %w",
			err,
		)
	}

	if _, err := db.ExecContext(
		ctx,
		"PRAGMA foreign_keys = ON;",
	); err != nil {
		_ = db.Close()

		return nil, fmt.Errorf(
			"erro ao habilitar chaves estrangeiras: %w",
			err,
		)
	}

	return db, nil
}
