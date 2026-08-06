package app

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func Nova(db *sql.DB) *App {
	router := gin.Default()

	container := novoContainer(db)

	registrarRotas(router, container)

	return &App{
		router: router,
	}
}

func (a *App) Executar(porta string) error {
	endereco := ":" + porta

	return a.router.Run(endereco)
}
