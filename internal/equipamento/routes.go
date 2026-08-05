package equipamento

import "github.com/gin-gonic/gin"

func RegistrarRotas(router *gin.RouterGroup, handler *Handler) {
	group := router.Group("/equipamentos")

	group.GET("", handler.Listar)
	group.GET("/codigo/:codigo", handler.BuscarPorCodigo)
	group.GET("/:id/modificacoes", handler.ListarModificacoes)
	group.GET("/:id", handler.BuscarPorID)
}
