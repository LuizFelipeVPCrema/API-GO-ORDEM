package personagem

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NovoRepositorio(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// func (r *Repository) Criar() error {}

// func (r *Repository) Listar() ([]Personagem, error) {}

// func (r *Repository) BuscarPorID() error {}

// func (r *Repository) Atualizar() error {}

// func (r *Repository) Deletar() error {}
