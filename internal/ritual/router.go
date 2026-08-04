package ritual

import "github.com/gin-gonic/gin"

func RegistrarRotas(router *gin.RouterGroup, handler *Handler) {
	grupo := router.Group("/rituais")

	grupo.GET("", handler.Listar)
	grupo.GET("/:id", handler.BuscarPorID)
	grupo.GET("/codigo/:codigo", handler.BuscarPorCodigo)
	grupo.GET("/:id/aprimoramentos", handler.ListarAprimoramentos)
}
