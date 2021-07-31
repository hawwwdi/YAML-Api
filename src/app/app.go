package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hawwwdi/YAML-Api/src/controllers"
)

type App struct {
	e *gin.Engine
}

func NewApp(r renderer) *App {
	e := gin.Default()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	e.POST("/template/:id", controllers.GetTemplateRenderer(r))
	return &App{
		e: e,
	}
}

func (a *App) Start(port string) error {
	return a.e.Run("localhost:" + port)
}
