package personagem

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type Repository interface {
	Listar(ctx context.Context, filtro Filtro) ([]Personagem, error)
	BuscarPorID(ctx context.Context, id int64) (*Personagem, error)
	Criar(ctx context.Context, personagem Personagem) (*Personagem, error)
	Atualizar(ctx context.Context, personagem Personagem) error
	Desativar(ctx context.Context, id int64) error
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

const selecaoPersonagem = `
	SELECT
		p.id,
		p.nome,
		p.jogador,
		p.classe_id,
		p.nex,
		p.pontos_prestigio,
		p.idade,
		p.aparencia,
		p.personalidade,
		p.historia,
		p.objetivo,
		p.ativa,

		a.agilidade,
		a.forca,
		a.intelecto,
		a.presenca,
		a.vigor,

		r.pv_atual,
		r.pv_maximo,
		r.pe_atual,
		r.pe_maximo,
		r.sanidade_atual,
		r.sanidade_maxima
	FROM personagens p
	INNER JOIN personagem_atributos a
		ON a.personagem_id = p.id
	INNER JOIN personagem_recursos r
		ON r.personagem_id = p.id
`

func (r *SQLiteRepository) Listar(ctx context.Context, filtro Filtro) ([]Personagem, error) {
	var query strings.Builder

	query.WriteString(selecaoPersonagem)
	query.WriteString(" WHERE p.ativa = 1")

	args := make([]any, 0, 4)

	if filtro.Nome != "" {
		query.WriteString(" AND LOWER(p.nome) LIKE LOWER(?)")

		args = append(args, "%"+filtro.Nome+"%")
	}

	if filtro.ClasseID != nil {
		query.WriteString(" AND p.classe_id = ? ")

		args = append(args, *filtro.ClasseID)
	}

	query.WriteString(`
		ORDER BY p.nome, p.id
		LIMIT ?
		OFFSET ?
	`)

	args = append(args, filtro.Limit, filtro.Offset)

	rows, err := r.db.QueryContext(ctx, query.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar personagens: %w", err)
	}
	defer rows.Close()

	personagens := make([]Personagem, 0)

	for rows.Next() {
		personagemEncontrado, err := escanearPersonagem(rows)

		if err != nil {
			return nil, fmt.Errorf("eroo ao ler personagens: %w", err)
		}

		personagens = append(personagens, personagemEncontrado)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante leitura dos personagens: %w", err)
	}

	return personagens, nil
}

func (r *SQLiteRepository) BuscarPorID(ctx context.Context, id int64) (*Personagem, error) {
	query := selecaoPersonagem + `
		WHERE p.id = ?
		AND p.ativa = 1
	`

	row := r.db.QueryRowContext(ctx, query, id)

	personagemEncontrado, err := escanearPersonagem(row)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrPersonagemNaoEncontrado
	}

	if err != nil {
		return nil, fmt.Errorf("erro ao consultar personagem por id: %w", err)
	}

	return &personagemEncontrado, nil
}

func (r *SQLiteRepository) Criar(ctx context.Context, personagem Personagem) (*Personagem, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao iniciar criacao do personagem: %w", err)
	}
	defer tx.Rollback()

	resultado, err := tx.ExecContext(
		ctx,
		`
			INSERT INTO personagens (
				nome,
				jogador,
				classe_id,
				nex,
				pontos_prestigio,
				idade,
				aparencia,
				personalidade,
				historia,
				objetivo,
				ativa
			)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 1)
		`,
		personagem.Nome,
		personagem.Jogador,
		personagem.ClasseID,
		personagem.NEX,
		personagem.PontosPrestigio,
		personagem.Idade,
		personagem.Aparencia,
		personagem.Personalidade,
		personagem.Historia,
		personagem.Objetivo,
	)
	if err != nil {
		return nil, fmt.Errorf("erro ao inserir personagem: %w", err)
	}

	personagem.ID, err = resultado.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("erro ao obter id do personagem: %w", err)
	}

	_, err = tx.ExecContext(
		ctx,
		`
			INSERT INTO personagem_atributos (
				personagem_id,
				agilidade,
				forca,
				intelecto,
				presenca,
				vigor
			)
			VALUES (?, ?, ?, ?, ?, ?)	
		`,
		personagem.ID,
		personagem.Atributos.Agilidade,
		personagem.Atributos.Forca,
		personagem.Atributos.Intelecto,
		personagem.Atributos.Presenca,
		personagem.Atributos.Vigor,
	)
	if err != nil {
		return nil, fmt.Errorf("erro ao inserir atributos: %w", err)
	}

	_, err = tx.ExecContext(
		ctx,
		`
			INSERT INTO personagem_recursos (
				personagem_id,
				pv_atual,
				pv_maximo,
				pe_atual,
				pe_maximo,
				sanidade_atual,
				sanidade_maxima
			)
			VALUES (?, ?, ?, ?, ?, ?, ?)
		`,
		personagem.ID,
		personagem.Recursos.PVAtual,
		personagem.Recursos.PVMaximo,
		personagem.Recursos.PEAtual,
		personagem.Recursos.PEMaximo,
		personagem.Recursos.SanidadeAtual,
		personagem.Recursos.SanidadeMaxima,
	)
	if err != nil {
		return nil, fmt.Errorf("erro ao inserir recursos: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("erro ao confirmar criacao do personagem: %w", err)
	}

	personagem.Ativa = true

	return &personagem, nil
}

func (r *SQLiteRepository) Atualizar(ctx context.Context, personagem Personagem) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("erro ao iniciar atualizacao: %w", err)
	}
	defer tx.Rollback()

	resultado, err := tx.ExecContext(
		ctx,
		`
			UPDATE personagens
			SET
				nome = ?,
				jogador = ?,
				classe_id = ?,
				nex = ?,
				pontos_prestigio = ?,
				idade = ?,
				aparencia = ?,
				personalidade = ?,
				historia = ?,
				objetivo = ?,
				atualizado_em = CURRENT_TIMESTAMP
			WHERE 
				id = ?
				AND ativa = 1
		`,
		personagem.Nome,
		personagem.Jogador,
		personagem.ClasseID,
		personagem.NEX,
		personagem.PontosPrestigio,
		personagem.Idade,
		personagem.Aparencia,
		personagem.Personalidade,
		personagem.Historia,
		personagem.Objetivo,
		personagem.ID,
	)
	if err != nil {
		return fmt.Errorf("erro ao atualizar personagem: %w", err)
	}

	afetadas, err := resultado.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar atualizacao: %w", err)
	}

	if afetadas == 0 {
		return ErrPersonagemNaoEncontrado
	}

	_, err = tx.ExecContext(
		ctx,
		`
			UPDATE personagem_atributos
			SET
				agilidade = ?,
				forca = ?,
				intelecto = ?,
				presenca = ?,
				vigor = ?
			WHERE personagem_id = ?
		`,
		personagem.Atributos.Agilidade,
		personagem.Atributos.Forca,
		personagem.Atributos.Intelecto,
		personagem.Atributos.Presenca,
		personagem.Atributos.Vigor,
		personagem.ID,
	)
	if err != nil {
		return fmt.Errorf("erro ao atualizar atributos: %w", err)
	}

	_, err = tx.ExecContext(
		ctx,
		`
			UPDATE personagem_recursos
			SET
				pv_atual = ?,
				pv_maximo = ?,
				pe_atual = ?,
				pe_maximo = ?,
				sanidade_atual = ?,
				sanidade_maxima = ?,
				atualizado_em = CURRENT_TIMESTAMP
			WHERE personagem_id = ?
		`,
		personagem.Recursos.PVAtual,
		personagem.Recursos.PVMaximo,
		personagem.Recursos.PEAtual,
		personagem.Recursos.PEMaximo,
		personagem.Recursos.SanidadeAtual,
		personagem.Recursos.SanidadeMaxima,
		personagem.ID,
	)
	if err != nil {
		return fmt.Errorf("erro ao atualizar recursos: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("erro ao confirmar atualizacao: %w", err)
	}

	return nil
}

func (r *SQLiteRepository) Desativar(ctx context.Context, id int64) error {
	resultado, err := r.db.ExecContext(
		ctx,
		`
			UPDATE personagens
			SET
				ativa = 0,
				atualizado_em = CURRENT_TIMESTAMP
			WHERE
				id = ?
				AND ativa = 1
		`,
		id,
	)
	if err != nil {
		return fmt.Errorf("erro ao remover personagem: %w", err)
	}

	afetadas, err := resultado.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar remocao: %w", err)
	}

	if afetadas == 0 {
		return ErrPersonagemNaoEncontrado
	}

	return nil
}

type scanner interface {
	Scan(dest ...any) error
}

func escanearPersonagem(row scanner) (Personagem, error) {
	var (
		personagemEncontrado Personagem

		jogador       sql.NullString
		idade         sql.NullInt64
		aparencia     sql.NullString
		personalidade sql.NullString
		historia      sql.NullString
		objetivo      sql.NullString

		ativo int
	)

	err := row.Scan(
		&personagemEncontrado.ID,
		&personagemEncontrado.Nome,
		&jogador,
		&personagemEncontrado.ClasseID,
		&personagemEncontrado.NEX,
		&personagemEncontrado.PontosPrestigio,
		&idade,
		&aparencia,
		&personalidade,
		&historia,
		&objetivo,
		&ativo,

		&personagemEncontrado.Atributos.Agilidade,
		&personagemEncontrado.Atributos.Forca,
		&personagemEncontrado.Atributos.Intelecto,
		&personagemEncontrado.Atributos.Presenca,
		&personagemEncontrado.Atributos.Vigor,

		&personagemEncontrado.Recursos.PVAtual,
		&personagemEncontrado.Recursos.PVMaximo,
		&personagemEncontrado.Recursos.PEAtual,
		&personagemEncontrado.Recursos.PEMaximo,
		&personagemEncontrado.Recursos.SanidadeAtual,
		&personagemEncontrado.Recursos.SanidadeMaxima,
	)
	if err != nil {
		return Personagem{}, err
	}

	personagemEncontrado.Ativa = ativo == 1

	if jogador.Valid {
		valor := jogador.String
		personagemEncontrado.Jogador = &valor
	}

	if idade.Valid {
		valor := int(idade.Int64)
		personagemEncontrado.Idade = &valor
	}

	if aparencia.Valid {
		valor := aparencia.String
		personagemEncontrado.Aparencia = &valor
	}

	if personalidade.Valid {
		valor := personalidade.String
		personagemEncontrado.Personalidade = &valor
	}

	if historia.Valid {
		valor := historia.String
		personagemEncontrado.Historia = &valor
	}

	if objetivo.Valid {
		valor := objetivo.String
		personagemEncontrado.Objetivo = &valor
	}

	return personagemEncontrado, nil
}
