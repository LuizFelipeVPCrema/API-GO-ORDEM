package main

import (
	"fmt"
	"log"

	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/app"
	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/config"
	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/database"
)

func main() {
	if err := executar(); err != nil {
		log.Fatal(err)
	}
}

func executar() error {
	cfg, err := config.Carregar()
	if err != nil {
		log.Fatalf("erro ao carregar configuracao: %v", err)
	}

	db, err := database.ConectarSQLite(cfg.Database.Path)
	if err != nil {
		log.Fatalf("erro ao conectar ao banco de dados: %v", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Printf(
				"erro ao fechar conexao com o banco: %v",
				err,
			)
		}
	}()
	if err := database.ExecutarMigrations(db); err != nil {
		log.Fatal(
			"erro ao executar migrations: ",
			err,
		)
	}

	log.Printf("conexao com SQLite realizada: &s", cfg.Database.Path)

	aplicacao := app.Nova(db)

	log.Printf("servidor iniciado em http://localhost:%s", cfg.Server.Port)

	if err := aplicacao.Executar(cfg.Server.Port); err != nil {
		return fmt.Errorf("erro ao iniciar servidor: %w", err)
	}

	return nil
}
