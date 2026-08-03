package personagem

import (
	"github.com/gin-gonic/gin"
)

func PersonagemRoutes(router *gin.Engine) {
	personagemRoutes := router.Group("/personagens")
	personagemRoutes.GET("/", GetPersonagem)
	personagemRoutes.POST("/", CreatePersonagem)
	personagemRoutes.PUT("/:id", UpdatePersonagem)
	personagemRoutes.DELETE("/:id", DeletePersonagem)
}
