package pericia

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type scanner interface {
	Scan(dest ...any) error
}

func escanearPericia(row scanner) (Pericia, error) {
	var (
		periciaEncontrada Pericia

		somenteTreinada       int
		penalidadeCarga       int
		permiteEspecializacao int
		ativa                 int
	)

	err := row.Scan(
		&periciaEncontrada.ID,
		&periciaEncontrada.Codigo,
		&periciaEncontrada.Nome,
		&periciaEncontrada.AtributoBase,
		&somenteTreinada,
		&penalidadeCarga,
		&permiteEspecializacao,
		&periciaEncontrada.OrdemExibicao,
		&ativa,
	)
	if err != nil {
		return Pericia{}, err
	}

	periciaEncontrada.SomenteTreinada = somenteTreinada == 1

	periciaEncontrada.PenalidadeCarga = penalidadeCarga == 1

	periciaEncontrada.PermiteEspecializacao = permiteEspecializacao == 1

	periciaEncontrada.Ativa = ativa == 1

	return periciaEncontrada, nil
}

type Repository interface {
	Listar(ctx context.Context) ([]Pericia, error)

	BuscarPorID(ctx context.Context, id int64) (*Pericia, error)

	BuscarPorCodigo(ctx context.Context, codigo Codigo) (*Pericia, error)
}

type SQLiteRepository struct {
	db *sql.DB
}

var _ Repository = (*SQLiteRepository)(nil)

func NovoRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{db: db}
}

func (r *SQLiteRepository) Listar(ctx context.Context) ([]Pericia, error) {
	const query = `
	    SELECT
			id,
			codigo,
			nome,
			atributo_base,
			somente_treinada,
			penalidade_carga,
			permite_especializacao,
			ordem_exibicao,
			ativa
		FROM pericias
		WHERE ativa = 1
		ORDER BY ordem_exibicao
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar pericias: %w", err)
	}
	defer rows.Close()

	pericias := make([]Pericia, 0)

	for rows.Next() {
		periciaEncontrada, err := escanearPericia(rows)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler pericia: %w", err)
		}

		pericias = append(pericias, periciaEncontrada)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante leitura das pericias: %w", err)
	}

	return pericias, nil
}

func (r *SQLiteRepository) BuscarPorID(ctx context.Context, id int64) (*Pericia, error) {
	const query = `
		SELECT
			id,
			codigo,
			nome,
			atributo_base,
			somente_treinada,
			penalidade_carga,
			permite_especializacao,
			ordem_exibicao,
			ativa
		FROM pericias
		WHERE
			id = ?
			AND ativa = 1
	`

	row := r.db.QueryRowContext(ctx, query, id)

	periciaEncontrada, err := escanearPericia(row)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrPericiaNaoEncontrada
	}

	if err != nil {
		return nil, fmt.Errorf("erro ao consultar pericia por id: %w", err)
	}

	return &periciaEncontrada, nil
}

func (r *SQLiteRepository) BuscarPorCodigo(ctx context.Context, codigo Codigo) (*Pericia, error) {
	const query = `
		SELECT
			id,
			codigo,
			nome,
			atributo_base,
			somente_treinada,
			penalidade_carga,
			permite_especializacao,
			ordem_exibicao,
			ativa
		FROM pericias
		WHERE
			codigo = ?
			AND ativa = 1
	`

	row := r.db.QueryRowContext(ctx, query, codigo)

	periciaEncontrada, err := escanearPericia(row)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrPericiaNaoEncontrada
	}

	if err != nil {
		return nil, fmt.Errorf("erro ao consultar pericia por codigo: %w", err)
	}

	return &periciaEncontrada, nil
}
