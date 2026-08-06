package main

import (
	"log"

	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/classe"
	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/config"
	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/database"
	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/equipamento"
	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/habilidade"
	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/patente"
	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/pericia"
	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/personagem"
	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/ritual"
	"github.com/gin-gonic/gin"
)

func main() {
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

	patenteRepository := patente.NovoRepository(db)
	patenteService := patente.NovoService(patenteRepository)
	patenteHandler := patente.NovoHandler(patenteService)

	periciaRepository := pericia.NovoRepository(db)
	periciaService := pericia.NovoService(periciaRepository)
	periciaHandler := pericia.NovoHandler(periciaService)

	classeRepository := classe.NovoRepository(db)
	classeService := classe.NovoService(classeRepository)
	classeHandler := classe.NovoHandler(classeService)

	habilidadeRepository := habilidade.NovoRepository(db)
	habilidadeService := habilidade.NovoService(habilidadeRepository)
	habilidadeHandler := habilidade.NovoHandler(habilidadeService)

	ritualRepository := ritual.NovoRepository(db)
	ritualService := ritual.NovoService(ritualRepository)
	ritualHandler := ritual.NovoHandler(ritualService)

	equipamentoRepository := equipamento.NovoRepository(db)
	equipamentoService := equipamento.NovoService(equipamentoRepository)
	equipamentoHandler := equipamento.NovoHandler(equipamentoService)

	personagemRepository := personagem.NovoRepository(db)
	personagemService := personagem.NovoService(personagemRepository, classeService)
	personagemHandler := personagem.NovoHandler(personagemService)

	router := gin.Default()

	api := router.Group("/api/v1")

	patente.RegistrarRotas(api, patenteHandler)
	pericia.RegistrarRota(api, periciaHandler)
	classe.RegistrarRotas(api, classeHandler)
	habilidade.RegisterRotas(api, habilidadeHandler)
	ritual.RegistrarRotas(api, ritualHandler)
	equipamento.RegistrarRotas(api, equipamentoHandler)
	personagem.RegistrarRotas(api, personagemHandler)

	enderecoServidor := ":" + cfg.Server.Port

	log.Printf(
		"servidor iniciado em http://localhost:%s",
		enderecoServidor,
	)

	if err := router.Run(enderecoServidor); err != nil {
		log.Fatalf("erro ao iniciar servidor HTTP: %v", err)
	}
}
