package habilidade

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type Repository interface {
	Listar(ctx context.Context) ([]Habilidade, error)
	BuscarPorID(ctx context.Context, id int64) (*Habilidade, error)
	BuscarPorCodigo(ctx context.Context, codigo Codigo) (*Habilidade, error)
	ListarPorClasseID(ctx context.Context, classeID int64) ([]HabilidadeClasse, error)
}

type SQLiteRepository struct {
	db *sql.DB
}

var _ Repository = (*SQLiteRepository)(nil)

func NovoRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{db: db}
}

func (r *SQLiteRepository) Listar(ctx context.Context) ([]Habilidade, error) {
	const query = `
		SELECT
			id,
			codigo,
			nome,
			categoria,
			tipo_ativacao,
			custo_pe_base,
			custo_pe_variavel,
			descricao_resumida,
			fonte_regra,
			versao_regra,
			ativa
		FROM habilidades
		WHERE ativa = 1
		ORDER BY nome
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar habilidades: %w", err)
	}
	defer rows.Close()

	habilidades := make([]Habilidade, 0)

	for rows.Next() {
		habilidadeEncontrada, err := escanearHabilidade(rows)

		if err != nil {
			return nil, fmt.Errorf("erro ao ler habilidade: %w", err)
		}

		habilidades = append(habilidades, habilidadeEncontrada)

	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante leitura das habilidades: %w", err)
	}

	return habilidades, nil

}

func (r *SQLiteRepository) BuscarPorID(ctx context.Context, id int64) (*Habilidade, error) {
	const query = `
		SELECT
			id,
			codigo,
			nome,
			categoria,
			tipo_ativacao,
			custo_pe_base,
			custo_pe_variavel,
			descricao_resumida,
			fonte_regra,
			versao_regra,
			ativa
		FROM habilidades
		WHERE 
			id = ?
			AND ativa = 1
	`

	row := r.db.QueryRowContext(ctx, query, id)

	habilidadeEncontrada, err := escanearHabilidade(row)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrHabilidadeNaoEncontrada
	}

	if err != nil {
		return nil, fmt.Errorf("erro ao consultar habilidade por id: %w", err)
	}

	return &habilidadeEncontrada, nil
}

func (r *SQLiteRepository) BuscarPorCodigo(ctx context.Context, codigo Codigo) (*Habilidade, error) {
	const query = `
		SELECT
			id,
			codigo,
			nome,
			categoria,
			tipo_ativacao,
			custo_pe_base,
			custo_pe_variavel,
			descricao_resumida,
			fonte_regra,
			versao_regra,
			ativa
		FROM habilidades
		WHERE
			codigo = ?
			AND ativa = 1
	`

	row := r.db.QueryRowContext(ctx, query, string(codigo))

	habilidadeEncontrada, err := escanearHabilidade(row)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrHabilidadeNaoEncontrada
	}

	if err != nil {
		return nil, fmt.Errorf("erro ao consultar habilidade por codigo: %w", err)
	}

	return &habilidadeEncontrada, nil

}

func (r *SQLiteRepository) ListarPorClasseID(ctx context.Context, classeID int64) ([]HabilidadeClasse, error) {
	var classeExiste bool

	err := r.db.QueryRowContext(
		ctx,
		`
			SELECT EXISTS(
				SELECT 1 
				FROM classes 
				WHERE 
					id = ?
					AND ativa = 1
			)
		`,
		classeID,
	).Scan(&classeExiste)

	if err != nil {
		return nil, fmt.Errorf("erro ao verificar classe: %w", err)
	}

	if !classeExiste {
		return nil, ErrClasseNaoEncontrada
	}

	const query = `
		SELECT
			ch.id,
			ch.classe_id,
			ch.nex_minimo,
			ch.forma_aquisicao,
			ch.ordem_exibicao,

			h.id,
			h.codigo,
			h.nome,
			h.categoria,
			h.tipo_ativacao,
			h.custo_pe_base,
			h.custo_pe_variavel,
			h.descricao_resumida,
			h.fonte_regra,
			h.versao_regra,
			h.ativa
		FROM classe_habilidades ch
		INNER JOIN habilidades h
			ON h.id = ch.habilidade_id
		WHERE
			ch.classe_id = ?
			AND h.ativa = 1
		ORDER BY
			ch.nex_minimo,
			ch.ordem_exibicao,
			h.nome
	`

	rows, err := r.db.QueryContext(
		ctx,
		query,
		classeID,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"erro ao consultar habilidades da classe: %w",
			err,
		)
	}
	defer rows.Close()

	vinculos := make([]HabilidadeClasse, 0)

	for rows.Next() {
		vinculo, err := escanearHabilidadeClasse(rows)
		if err != nil {
			return nil, fmt.Errorf(
				"erro ao ler habilidade da classe: %w",
				err,
			)
		}

		vinculos = append(vinculos, vinculo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"erro durante leitura das habilidades da classe: %w",
			err,
		)
	}

	return vinculos, nil
}

type scanner interface {
	Scan(dest ...any) error
}

func escanearHabilidade(
	row scanner,
) (Habilidade, error) {
	var (
		habilidadeEncontrada Habilidade

		codigo        string
		categoria     string
		tipoAtivacao  string
		custoPEBase   sql.NullInt64
		custoVariavel int
		descricao     sql.NullString
		versaoRegra   sql.NullString
		ativa         int
	)

	err := row.Scan(
		&habilidadeEncontrada.ID,
		&codigo,
		&habilidadeEncontrada.Nome,
		&categoria,
		&tipoAtivacao,
		&custoPEBase,
		&custoVariavel,
		&descricao,
		&habilidadeEncontrada.FonteRegra,
		&versaoRegra,
		&ativa,
	)
	if err != nil {
		return Habilidade{}, err
	}

	habilidadeEncontrada.Codigo = Codigo(codigo)
	habilidadeEncontrada.Categoria = Categoria(categoria)
	habilidadeEncontrada.TipoAtivacao =
		TipoAtivacao(tipoAtivacao)

	habilidadeEncontrada.CustoPEVariavel =
		custoVariavel == 1

	habilidadeEncontrada.Ativa = ativa == 1

	if custoPEBase.Valid {
		valor := int(custoPEBase.Int64)
		habilidadeEncontrada.CustoPEBase = &valor
	}

	if descricao.Valid {
		valor := descricao.String
		habilidadeEncontrada.DescricaoResumida = &valor
	}

	if versaoRegra.Valid {
		valor := versaoRegra.String
		habilidadeEncontrada.VersaoRegra = &valor
	}

	return habilidadeEncontrada, nil
}

func escanearHabilidadeClasse(
	row scanner,
) (HabilidadeClasse, error) {
	var (
		vinculo HabilidadeClasse

		codigo         string
		categoria      string
		tipoAtivacao   string
		formaAquisicao string

		custoPEBase   sql.NullInt64
		custoVariavel int
		descricao     sql.NullString
		versaoRegra   sql.NullString
		ativa         int
	)

	err := row.Scan(
		&vinculo.ID,
		&vinculo.ClasseID,
		&vinculo.NEXMinimo,
		&formaAquisicao,
		&vinculo.OrdemExibicao,

		&vinculo.Habilidade.ID,
		&codigo,
		&vinculo.Habilidade.Nome,
		&categoria,
		&tipoAtivacao,
		&custoPEBase,
		&custoVariavel,
		&descricao,
		&vinculo.Habilidade.FonteRegra,
		&versaoRegra,
		&ativa,
	)
	if err != nil {
		return HabilidadeClasse{}, err
	}

	vinculo.Aquisicao =
		FormaAquisicao(formaAquisicao)

	vinculo.Habilidade.Codigo = Codigo(codigo)
	vinculo.Habilidade.Categoria = Categoria(categoria)
	vinculo.Habilidade.TipoAtivacao =
		TipoAtivacao(tipoAtivacao)

	vinculo.Habilidade.CustoPEVariavel =
		custoVariavel == 1

	vinculo.Habilidade.Ativa = ativa == 1

	if custoPEBase.Valid {
		valor := int(custoPEBase.Int64)
		vinculo.Habilidade.CustoPEBase = &valor
	}

	if descricao.Valid {
		valor := descricao.String
		vinculo.Habilidade.DescricaoResumida = &valor
	}

	if versaoRegra.Valid {
		valor := versaoRegra.String
		vinculo.Habilidade.VersaoRegra = &valor
	}

	return vinculo, nil
}
