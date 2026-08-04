package ritual

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type Repository interface {
	Listar(ctx context.Context, filtro Filtro) ([]Ritual, error)
	BuscarPorID(ctx context.Context, id int64) (*Ritual, error)
	BuscarPorCodigo(ctx context.Context, codigo Codigo) (*Ritual, error)
	ListarAprimoramentos(ctx context.Context, ritualID int64) ([]Aprimoramento, error)
}

type SQLiteRepository struct {
	db *sql.DB
}

var _ Repository = (*SQLiteRepository)(nil)

func NovoRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{db: db}
}

func (r *SQLiteRepository) Listar(ctx context.Context, filtro Filtro) ([]Ritual, error) {
	query := `
		SELECT
			id,
			codigo,
			nome,
			elemento,
			circulo,
			execucao,
			alcance,
			alvo,
			area,
			duracao,
			resistencia,
			custo_pe_base,
			requer_componente,
			descricao_resumida,
			fonte_regra,
			versao_regra,
			pagina_referencia,
			ativo
		FROM rituais
		WHERE ativo = 1
	`

	args := make([]any, 0, 2)

	if filtro.Elemento != nil {
		query += " AND elemento = ?"
		args = append(args, string(*filtro.Elemento))
	}

	if filtro.Circulo != nil {
		query += " AND circulo = ?"
		args = append(args, int(*filtro.Circulo))
	}

	query += " ORDER BY circulo, elemento, nome"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar rituais: %w", err)
	}
	defer rows.Close()

	rituais := make([]Ritual, 0)

	for rows.Next() {
		ritualEncontrado, err := escanearRitual(rows)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler ritual: %w", err)
		}
		rituais = append(rituais, ritualEncontrado)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante leitura dos rituais: %w", err)
	}

	return rituais, nil
}

func (r *SQLiteRepository) BuscarPorID(ctx context.Context, id int64) (*Ritual, error) {
	const query = `
		SELECT
			id,
			codigo,
			nome,
			elemento,
			circulo,
			execucao,
			alcance,
			alvo,
			area,
			duracao,
			resistencia,
			custo_pe_base,
			requer_componente,
			descricao_resumida,
			fonte_regra,
			versao_regra,
			pagina_referencia,
			ativo
		FROM rituais
		WHERE
			id = ?
			AND ativo = 1
	`

	row := r.db.QueryRowContext(ctx, query, id)

	ritualEncontrado, err := escanearRitual(row)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrRitualNaoEncontrado
	}

	if err != nil {
		return nil, fmt.Errorf("erro ao consultar ritual por id: %w", err)
	}

	return &ritualEncontrado, nil
}

func (r *SQLiteRepository) BuscarPorCodigo(ctx context.Context, codigo Codigo) (*Ritual, error) {
	const query = `
		SELECT
			id,
			codigo,
			nome,
			elemento,
			circulo,
			execucao,
			alcance,
			alvo,
			area,
			duracao,
			resistencia,
			custo_pe_base,
			requer_componente,
			descricao_resumida,
			fonte_regra,
			versao_regra,
			pagina_referencia,
			ativo
		FROM rituais
		WHERE
			codigo = ?
			AND ativo = 1
	`

	row := r.db.QueryRowContext(ctx, query, string(codigo))

	ritualEncontrado, err := escanearRitual(row)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrRitualNaoEncontrado
	}

	if err != nil {
		return nil, fmt.Errorf("erro ao consultar ritual por codigo: %w", err)
	}

	return &ritualEncontrado, nil
}

func (r *SQLiteRepository) ListarAprimoramentos(ctx context.Context, ritualID int64) ([]Aprimoramento, error) {
	const query = `
		SELECT
			id,
			ritual_id,
			tipo,
			custo_pe_adicional,
			nex_minimo,
			circulo_minimo,
			descricao_resumida,
			ordem_exibicao,
			ativo
		FROM ritual_aprimoramentos
		WHERE
			ritual_id = ?
			AND ativo = 1
		ORDER BY ordem_exibicao
	`

	rows, err := r.db.QueryContext(ctx, query, ritualID)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar aprimoramentos do ritual: %w", err)
	}
	defer rows.Close()

	aprimoramentos := make([]Aprimoramento, 0)

	for rows.Next() {
		aprimoramento, err := escanearAprimoramento(rows)

		if err != nil {
			return nil, fmt.Errorf("erro ao ler aprimoramento: %w", err)
		}

		aprimoramentos = append(aprimoramentos, aprimoramento)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante leitura dos aprimoramentos: %w", err)
	}

	return aprimoramentos, nil
}

type scanner interface {
	Scan(dest ...any) error
}

func escanearRitual(row scanner) (Ritual, error) {
	var (
		ritualEncontrado Ritual

		codigo   string
		elemento string
		circulo  int
		execucao string

		alvo              sql.NullString
		area              sql.NullString
		resistencia       sql.NullString
		descricaoResumida sql.NullString
		versaoRegra       sql.NullString
		paginaReferencia  sql.NullInt64

		requerComponente int
		ativo            int
	)

	err := row.Scan(
		&ritualEncontrado.ID,
		&codigo,
		&ritualEncontrado.Nome,
		&elemento,
		&circulo,
		&execucao,
		&ritualEncontrado.Alcance,
		&alvo,
		&area,
		&ritualEncontrado.Duracao,
		&resistencia,
		&ritualEncontrado.CustoPEBase,
		&requerComponente,
		&descricaoResumida,
		&ritualEncontrado.FonteRegra,
		&versaoRegra,
		&paginaReferencia,
		&ativo,
	)
	if err != nil {
		return Ritual{}, err
	}

	ritualEncontrado.Codigo = Codigo(codigo)
	ritualEncontrado.Elemento = Elemento(elemento)
	ritualEncontrado.Circulo = Circulo(circulo)
	ritualEncontrado.Execucao = TipoExecucao(execucao)

	ritualEncontrado.RequerComponente =
		requerComponente == 1

	ritualEncontrado.Ativo = ativo == 1

	if alvo.Valid {
		valor := alvo.String
		ritualEncontrado.Alvo = &valor
	}

	if area.Valid {
		valor := area.String
		ritualEncontrado.Area = &valor
	}

	if resistencia.Valid {
		valor := resistencia.String
		ritualEncontrado.Resistencia = &valor
	}

	if descricaoResumida.Valid {
		valor := descricaoResumida.String
		ritualEncontrado.DescricaoResumida = &valor
	}

	if versaoRegra.Valid {
		valor := versaoRegra.String
		ritualEncontrado.VersaoRegra = &valor
	}

	if paginaReferencia.Valid {
		valor := int(paginaReferencia.Int64)
		ritualEncontrado.PaginaReferencia = &valor
	}

	return ritualEncontrado, nil
}

func escanearAprimoramento(row scanner) (Aprimoramento, error) {
	var (
		aprimoramento Aprimoramento
		tipo          string

		nexMinimo         sql.NullInt64
		circuloMinimo     sql.NullInt64
		descricaoResumida sql.NullString

		ativo int
	)

	err := row.Scan(
		&aprimoramento.ID,
		&aprimoramento.RitualID,
		&tipo,
		&aprimoramento.CustoPEAdicional,
		&nexMinimo,
		&circuloMinimo,
		&descricaoResumida,
		&aprimoramento.OrdemExibicao,
		&ativo,
	)
	if err != nil {
		return Aprimoramento{}, err
	}

	aprimoramento.Tipo = TipoAprimoramento(tipo)
	aprimoramento.Ativo = ativo == 1

	if nexMinimo.Valid {
		valor := int(nexMinimo.Int64)
		aprimoramento.NEXMinimo = &valor
	}

	if circuloMinimo.Valid {
		valor := Circulo(circuloMinimo.Int64)
		aprimoramento.CirculoMinimo = &valor
	}

	if descricaoResumida.Valid {
		valor := descricaoResumida.String
		aprimoramento.DescricaoResumida = &valor
	}

	return aprimoramento, nil
}
