package pericia

import "github.com/gin-gonic/gin"

func RegistrarRota(router *gin.RouterGroup, handler *Handler) {
	grupo := router.Group("/pericias")

	grupo.GET("/", handler.Listar)
	grupo.GET("/:id", handler.BuscarPorID)
	grupo.GET("/codigo/:codigo", handler.BuscarPorCodigo)
}
