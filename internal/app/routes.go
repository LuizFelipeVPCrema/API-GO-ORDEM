package app

import (
	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/classe"
	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/equipamento"
	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/habilidade"
	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/patente"
	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/pericia"
	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/personagem"
	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/ritual"
	"github.com/gin-gonic/gin"
)

func registrarRotas(router *gin.Engine, container *container) {
	api := router.Group("/api/v1")

	patente.RegistrarRotas(api, container.patenteHandler)
	pericia.RegistrarRota(api, container.periciaHandler)
	classe.RegistrarRotas(api, container.classeHandler)
	habilidade.RegisterRotas(api, container.habilidadeHandler)
	ritual.RegistrarRotas(api, container.ritualHandler)
	equipamento.RegistrarRotas(api, container.equipamentoHandler)
	personagem.RegistrarRotas(api, container.personagemHandler)

}
