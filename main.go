package main

import (
	"prog2-2025-dic-Heinen/handlers"
	"prog2-2025-dic-Heinen/middlewares"
	"prog2-2025-dic-Heinen/services"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router := gin.Default()
	service := services.NewClimaService()
	handler := handlers.NewClimaHandler(service)
	mappingRoutes(router, handler)
	router.Run()
}

func mappingRoutes(router *gin.Engine, handler *handlers.ClimaHandler) {
	clima := router.Group("/clima").Use(middlewares.ValidateAuthHeader())
	{
		clima.POST("/metricas", handler.CalcularMetricas)
		clima.GET("/proyeccion", handler.CalcularProyeccion)
		clima.GET("/reporte", handler.ReporteOperaciones)
		clima.POST("/registro", handler.RegistroEstaciones)
	}
}
