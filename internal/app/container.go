package app

import (
	"database/sql"

	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/classe"
	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/equipamento"
	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/habilidade"
	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/patente"
	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/pericia"
	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/personagem"
	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/ritual"
)

type container struct {
	patenteHandler     *patente.Handler
	periciaHandler     *pericia.Handler
	classeHandler      *classe.Handler
	habilidadeHandler  *habilidade.Handler
	ritualHandler      *ritual.Handler
	equipamentoHandler *equipamento.Handler
	personagemHandler  *personagem.Handler
}

func novoContainer(db *sql.DB) *container {
	patenteHandler := novoPatenteHandler(db)
	periciaHandler := novoPericiaHandler(db)
	classeService, classeHandler := novoClasseHandler(db)
	habilidadeHandler := novoHabilidadeHandler(db)
	ritualHandler := novoRitualHandler(db)
	equipamentoHandler := novoEquipamentoHandler(db)
	personagemHandler := novoPersonagemHandler(db, classeService)

	return &container{
		patenteHandler:     patenteHandler,
		periciaHandler:     periciaHandler,
		classeHandler:      classeHandler,
		habilidadeHandler:  habilidadeHandler,
		ritualHandler:      ritualHandler,
		equipamentoHandler: equipamentoHandler,
		personagemHandler:  personagemHandler,
	}
}

func novoPatenteHandler(db *sql.DB) *patente.Handler {
	repository := patente.NovoRepository(db)
	service := patente.NovoService(repository)

	return patente.NovoHandler(service)
}

func novoPericiaHandler(db *sql.DB) *pericia.Handler {
	repository := pericia.NovoRepository(db)
	service := pericia.NovoService(repository)

	return pericia.NovoHandler(service)
}

func novoClasseHandler(db *sql.DB) (*classe.Service, *classe.Handler) {
	repository := classe.NovoRepository(db)
	service := classe.NovoService(repository)
	handler := classe.NovoHandler(service)

	return service, handler
}

func novoHabilidadeHandler(db *sql.DB) *habilidade.Handler {
	repository := habilidade.NovoRepository(db)
	service := habilidade.NovoService(repository)

	return habilidade.NovoHandler(service)
}

func novoRitualHandler(db *sql.DB) *ritual.Handler {
	repository := ritual.NovoRepository(db)
	service := ritual.NovoService(repository)

	return ritual.NovoHandler(service)
}

func novoEquipamentoHandler(db *sql.DB) *equipamento.Handler {
	repository := equipamento.NovoRepository(db)
	service := equipamento.NovoService(repository)

	return equipamento.NovoHandler(service)
}

func novoPersonagemHandler(db *sql.DB, classeService *classe.Service) *personagem.Handler {
	repository := personagem.NovoRepository(db)
	service := personagem.NovoService(repository, classeService)

	return personagem.NovoHandler(service)
}
