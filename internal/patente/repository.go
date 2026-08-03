package patente

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository interface {
	Listar(ctx context.Context) ([]Patente, error)

	BuscarPorID(ctx context.Context, id int64) (*Patente, error)

	BuscarPorPrestigio(ctx context.Context, pontos int) (*Patente, error)
}

type SQLiteRepository struct {
	db *sql.DB
}

func NovoRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (r *SQLiteRepository) Listar(ctx context.Context) ([]Patente, error) {
	const query = `
		SELECT
			p.id,
			p.codigo,
			p.nome,
			p.pontos_prestigio_minimos,
			p.limite_credito,
			p.nivel_hierarquico,
			p.ativa,
			l.categoria,
			l.quantidade_maxima
		FROM patentes p
		LEFT JOIN patente_limites_item l
			ON l.patente_id = p.id
		WHERE p.ativa = 1
		ORDER BY
			p.nivel_hierarquico,
			l.categoria
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf(
			"erro ao consultar patentes: %w",
			err,
		)
	}
	defer rows.Close()

	patentes, err := lerPatentes(rows)
	if err != nil {
		return nil, fmt.Errorf(
			"erro ao processar patentes: %w",
			err,
		)
	}

	return patentes, nil
}

func (r *SQLiteRepository) BuscarPorID(ctx context.Context, id int64) (*Patente, error) {
	const query = `
		SELECT
			p.id,
			p.codigo,
			p.nome,
			p.pontos_prestigio_minimos,
			p.limite_credito,
			p.nivel_hierarquico,
			p.ativa,
			l.categoria,
			l.quantidade_maxima
		FROM patentes p
		LEFT JOIN patente_limites_item l
			ON l.patente_id = p.id
		WHERE
			p.id = ?
			AND p.ativa = 1
		ORDER BY l.categoria
	`

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf(
			"erro ao consultar patente por id: %w",
			err,
		)
	}
	defer rows.Close()

	patentes, err := lerPatentes(rows)
	if err != nil {
		return nil, fmt.Errorf(
			"erro ao processar patente: %w",
			err,
		)
	}

	if len(patentes) == 0 {
		return nil, ErrPatenteNaoEncontrada
	}

	return &patentes[0], nil
}

func (r *SQLiteRepository) BuscarPorPrestigio(ctx context.Context, pontos int) (*Patente, error) {
	const query = `
		SELECT
			p.id,
			p.codigo,
			p.nome,
			p.pontos_prestigio_minimos,
			p.limite_credito,
			p.nivel_hierarquico,
			p.ativa,
			l.categoria,
			l.quantidade_maxima
		FROM patentes p
		LEFT JOIN patente_limites_item l
			ON l.patente_id = p.id
		WHERE p.id = (
			SELECT p2.id
			FROM patentes p2
			WHERE
				p2.ativa = 1
				AND p2.pontos_prestigio_minimos <= ?
			ORDER BY
				p2.pontos_prestigio_minimos DESC
			LIMIT 1
		)
		ORDER BY l.categoria
	`

	rows, err := r.db.QueryContext(
		ctx,
		query,
		pontos,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"erro ao consultar patente por pontos de prestígio: %w",
			err,
		)
	}
	defer rows.Close()

	patentes, err := lerPatentes(rows)
	if err != nil {
		return nil, fmt.Errorf(
			"erro ao processar patente por pontos de prestígio: %w",
			err,
		)
	}

	if len(patentes) == 0 {
		return nil, ErrPatenteNaoEncontrada
	}

	return &patentes[0], nil
}

func lerPatentes(rows *sql.Rows) ([]Patente, error) {
	patentes := make([]Patente, 0)

	indices := make(map[int64]int)

	for rows.Next() {
		var (
			id                     int64
			codigo                 string
			nome                   string
			pontosPrestigioMinimos int
			limiteCredito          string
			nivelHierarquico       int
			ativa                  int

			categoria        sql.NullInt64
			quantidadeMaxima sql.NullInt64
		)

		err := rows.Scan(
			&id,
			&codigo,
			&nome,
			&pontosPrestigioMinimos,
			&limiteCredito,
			&nivelHierarquico,
			&ativa,
			&categoria,
			&quantidadeMaxima,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"erro ao ler patente: %w",
				err,
			)
		}

		indice, encontrada := indices[id]

		if !encontrada {
			patente := Patente{
				ID:                     uint(id),
				Codigo:                 Codigo(codigo),
				Nome:                   nome,
				PontosPrestigioMinimos: pontosPrestigioMinimos,
				LimiteCredito:          LimiteCredito(limiteCredito),
				NivelHierarquico:       nivelHierarquico,
				Ativa:                  ativa == 1,
				Limites:                make([]LimiteItem, 0, 4),
			}

			patentes = append(patentes, patente)

			indice = len(patentes) - 1
			indices[id] = indice
		}

		if categoria.Valid && quantidadeMaxima.Valid {
			patentes[indice].Limites = append(
				patentes[indice].Limites,
				LimiteItem{
					Categoria: CategoriaItem(
						categoria.Int64,
					),
					QuantidadeMaxima: int(
						quantidadeMaxima.Int64,
					),
				},
			)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"erro durante leitura das patentes: %w",
			err,
		)
	}

	return patentes, nil
}
