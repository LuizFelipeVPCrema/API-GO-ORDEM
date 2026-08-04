package classe

import "github.com/gin-gonic/gin"

func RegistrarRotas(router *gin.RouterGroup, handler *Handler) {
	grupo := router.Group("/classes")

	grupo.GET("", handler.Listar)
	grupo.GET("/:id", handler.BuscarPorID)
	grupo.GET("/codigo/:codigo", handler.BuscarPorCodigo)
}
