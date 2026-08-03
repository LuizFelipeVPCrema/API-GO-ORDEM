package patente

import "github.com/gin-gonic/gin"

func RegistrarRotas(router *gin.RouterGroup, handler *Handler) {
	grupo := router.Group("/patentes")

	grupo.GET("", handler.Listar)
	grupo.GET("/por-prestigio/:pontos", handler.BuscarPorPrestigio)
	grupo.GET("/:id", handler.BuscarPorID)
}
