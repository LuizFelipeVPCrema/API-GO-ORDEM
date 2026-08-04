package habilidade

import "github.com/gin-gonic/gin"

func RegisterRotas(router *gin.RouterGroup, handler *Handler) {
	grupo := router.Group("/habilidades")

	grupo.GET("", handler.Listar)
	grupo.GET("/:id", handler.BuscarPorID)
	grupo.GET("/codigo/:codigo", handler.BuscarPorCodigo)
	grupo.GET("/classe/:id/habilidades", handler.ListarPorClasseID)
}
