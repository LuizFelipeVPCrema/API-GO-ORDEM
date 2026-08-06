package personagem

import (
	"github.com/gin-gonic/gin"
)

func RegistrarRotas(router *gin.RouterGroup, handler *Handler) {
	grupo := router.Group("/personagens")

	grupo.POST("", handler.Criar)
	grupo.GET("", handler.Listar)
	grupo.GET("/:id", handler.BuscarPorID)
	grupo.PATCH("/:id", handler.Atualizar)
	grupo.DELETE("/:id", handler.Remover)
}
