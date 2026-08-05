package equipamento

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type Repository interface {
	Listar(ctx context.Context, filtro Filtro) ([]Equipamento, error)
	BuscarPorID(ctx context.Context, id int64) (*Equipamento, error)
	BuscarPorCodigo(ctx context.Context, codigo Codigo) (*Equipamento, error)
	ListarModificacoes(ctx context.Context, equipamentoID int64) ([]Modificacao, error)
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

func (r *SQLiteRepository) Listar(ctx context.Context, filtro Filtro) ([]Equipamento, error) {
	query := `
		SELECT
			id,
			codigo,
			nome,
			tipo,
			categoria_base,
			espacos_base,
			descricao_resumida,
			fonte_regra,
			versao_regra,
			pagina_referencia,
			ativa
		FROM equipamentos
		WHERE ativa = 1
	`

	args := make([]any, 0, 2)

	if filtro.Tipo != nil {
		query += " AND tipo = ?"

		args = append(args, string(*filtro.Tipo))
	}

	if filtro.Categoria != nil {
		query += " AND categoria_base = ?"

		args = append(args, int(*filtro.Categoria))
	}

	query += `
		ORDER BY
			categoria_base,
			tipo,
			nome
	`

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar equipamento: %w", err)
	}
	defer rows.Close()

	equipamentos := make([]Equipamento, 0)

	for rows.Next() {
		equipamentoEncontrado, err := escanearEquipamento(rows)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler equipamento: %w", err)
		}

		equipamentos = append(equipamentos, equipamentoEncontrado)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante leitura dos equipamentos: %w", err)
	}

	return equipamentos, nil
}

func (r *SQLiteRepository) BuscarPorID(ctx context.Context, id int64) (*Equipamento, error) {
	const query = `
		SELECT
			id,
			codigo,
			nome,
			tipo,
			categoria_base,
			espacos_base,
			descricao_resumida,
			fonte_regra,
			versao_regra,
			pagina_referencia,
			ativa
		FROM equipamentos
		WHERE
			id = ?
			AND ativa = 1
	`

	row := r.db.QueryRowContext(ctx, query, id)

	equipamentoEncontrado, err := escanearEquipamento(row)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrEquipamentoNaoEncontrado
	}

	if err != nil {
		return nil, fmt.Errorf("erro ao consultar equipamento por id: %w", err)
	}

	return &equipamentoEncontrado, nil
}

func (r *SQLiteRepository) BuscarPorCodigo(ctx context.Context, codigo Codigo) (*Equipamento, error) {
	const query = `
		SELECT
			id,
			codigo,
			nome,
			tipo,
			categoria_base,
			espacos_base,
			descricao_resumida,
			fonte_regra,
			versao_regra,
			pagina_referencia,
			ativa
		FROM equipamentos
		WHERE
			codigo = ?
			AND ativa = 1
	`

	row := r.db.QueryRowContext(ctx, query, string(codigo))

	equipamentoEncontrado, err := escanearEquipamento(row)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrEquipamentoNaoEncontrado
	}

	if err != nil {
		return nil, fmt.Errorf("erro ao consultar equipamento por codigo: %w", err)
	}

	return &equipamentoEncontrado, nil
}

func (r *SQLiteRepository) ListarModificacoes(ctx context.Context, equipamentoID int64) ([]Modificacao, error) {
	const query = `
		SELECT DISTINCT
			m.id,
			m.codigo,
			m.nome,
			m.incremento_categoria,
			m.incremento_espacos,
			m.limite_por_item,
			m.cumulativa,
			m.descricao_resumida,
			m.fonte_regra,
			m.versao_regra,
			m.pagina_referencia,
			m.ativa
		FROM equipamentos e
		INNER JOIN modificacao_tipos_equipamento mt
			ON (
				mt.tipo_equipamento = e.tipo
				OR mt.tipo_equipamento = 'TODOS'
			)
		INNER JOIN modificacoes_equipamento m
			ON m.id = mt.modificacao_id
		WHERE
			e.id = ?
			AND e.ativa = 1
			AND m.ativa = 1
		ORDER BY m.nome
	`

	rows, err := r.db.QueryContext(ctx, query, equipamentoID)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar modificacoes do equipamento: %w", err)
	}
	defer rows.Close()

	modificacoes := make([]Modificacao, 0)

	for rows.Next() {
		modificacaoEncontrada, err := escanearModificacao(rows)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler modificacao: %w", err)
		}

		modificacoes = append(modificacoes, modificacaoEncontrada)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante leitura das modificacoes: %w", err)
	}

	return modificacoes, nil
}

type scanner interface {
	Scan(dest ...any) error
}

func escanearEquipamento(
	row scanner,
) (Equipamento, error) {
	var (
		equipamentoEncontrado Equipamento

		codigo    string
		tipo      string
		categoria int

		descricaoResumida sql.NullString
		versaoRegra       sql.NullString
		paginaReferencia  sql.NullInt64

		ativo int
	)

	err := row.Scan(
		&equipamentoEncontrado.ID,
		&codigo,
		&equipamentoEncontrado.Nome,
		&tipo,
		&categoria,
		&equipamentoEncontrado.EspacosBase,
		&descricaoResumida,
		&equipamentoEncontrado.FonteRegra,
		&versaoRegra,
		&paginaReferencia,
		&ativo,
	)
	if err != nil {
		return Equipamento{}, err
	}

	equipamentoEncontrado.Codigo =
		Codigo(codigo)

	equipamentoEncontrado.Tipo =
		Tipo(tipo)

	equipamentoEncontrado.CategoriaBase =
		Categoria(categoria)

	equipamentoEncontrado.Ativa =
		ativo == 1

	if descricaoResumida.Valid {
		valor := descricaoResumida.String

		equipamentoEncontrado.DescricaoResumida =
			&valor
	}

	if versaoRegra.Valid {
		valor := versaoRegra.String

		equipamentoEncontrado.VersaoRegra =
			&valor
	}

	if paginaReferencia.Valid {
		valor := int(
			paginaReferencia.Int64,
		)

		equipamentoEncontrado.PaginaReferencia =
			&valor
	}

	return equipamentoEncontrado, nil
}

func escanearModificacao(
	row scanner,
) (Modificacao, error) {
	var (
		modificacaoEncontrada Modificacao

		codigo string

		cumulativa int
		ativa      int

		descricaoResumida sql.NullString
		versaoRegra       sql.NullString
		paginaReferencia  sql.NullInt64
	)

	err := row.Scan(
		&modificacaoEncontrada.ID,
		&codigo,
		&modificacaoEncontrada.Nome,
		&modificacaoEncontrada.IncrementoCategoria,
		&modificacaoEncontrada.IncrementoEspacos,
		&modificacaoEncontrada.LimitePorItem,
		&cumulativa,
		&descricaoResumida,
		&modificacaoEncontrada.FonteRegra,
		&versaoRegra,
		&paginaReferencia,
		&ativa,
	)
	if err != nil {
		return Modificacao{}, err
	}

	modificacaoEncontrada.Codigo =
		Codigo(codigo)

	modificacaoEncontrada.Cumulativa =
		cumulativa == 1

	modificacaoEncontrada.Ativa =
		ativa == 1

	if descricaoResumida.Valid {
		valor := descricaoResumida.String

		modificacaoEncontrada.DescricaoResumida =
			&valor
	}

	if versaoRegra.Valid {
		valor := versaoRegra.String

		modificacaoEncontrada.VersaoRegra =
			&valor
	}

	if paginaReferencia.Valid {
		valor := int(
			paginaReferencia.Int64,
		)

		modificacaoEncontrada.PaginaReferencia =
			&valor
	}

	return modificacaoEncontrada, nil
}
