package classe

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type Repository interface {
	Listar(ctx context.Context) ([]Classe, error)

	BuscarPorID(ctx context.Context, id int64) (*Classe, error)

	BuscarPorCodigo(ctx context.Context, codigo Codigo) (*Classe, error)
}

type SQLiteRepository struct {
	db *sql.DB
}

var _ Repository = (*SQLiteRepository)(nil)

func NovoRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (r *SQLiteRepository) Listar(ctx context.Context) ([]Classe, error) {
	const query = `
		SELECT
			id,
			codigo,
			nome,
			ordem_exibicao,
			ativa
		FROM classes
		WHERE ativa = 1
		ORDER BY ordem_exibicao
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar classes: %w", err)
	}
	defer rows.Close()

	classes := make([]Classe, 0)

	for rows.Next() {
		classeEncontrada, err := escanearClasse(rows)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler classe: %w", err)
		}

		classes = append(classes, classeEncontrada)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante leitura das classes: %w", err)
	}

	return classes, nil
}

func (r *SQLiteRepository) BuscarPorID(ctx context.Context, id int64) (*Classe, error) {
	const query = `
		SELECT
			id,
			codigo,
			nome,
			ordem_exibicao,
			ativa
		FROM classes
		WHERE
			id = ?
			AND ativa = 1
	`

	row := r.db.QueryRowContext(ctx, query, id)

	classeEncontrada, err := escanearClasse(row)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrClasseNaoEncontrada
	}

	if err != nil {
		return nil, fmt.Errorf("erro ao consultar classe por id: %w", err)
	}

	return &classeEncontrada, nil
}

func (r *SQLiteRepository) BuscarPorCodigo(ctx context.Context, codigo Codigo) (*Classe, error) {
	const query = `
		SELECT
			id,
			codigo,
			nome,
			ordem_exibicao,
			ativa
		FROM classes
		WHERE
			codigo = ?
			AND ativa = 1
	`

	row := r.db.QueryRowContext(ctx, query, string(codigo))

	classeEncontrada, err := escanearClasse(row)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrClasseNaoEncontrada
	}

	if err != nil {
		return nil, fmt.Errorf("erro ao consultar classe por codigo: %w", err)
	}

	return &classeEncontrada, nil
}

type scanner interface {
	Scan(dest ...any) error
}

func escanearClasse(
	row scanner,
) (Classe, error) {
	var (
		classeEncontrada Classe
		codigo           string
		ativa            int
	)

	err := row.Scan(
		&classeEncontrada.ID,
		&codigo,
		&classeEncontrada.Nome,
		&classeEncontrada.OrdemExibicao,
		&ativa,
	)
	if err != nil {
		return Classe{}, err
	}

	classeEncontrada.Codigo = Codigo(codigo)
	classeEncontrada.Ativa = ativa == 1

	return classeEncontrada, nil
}
