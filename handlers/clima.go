package handlers

import (
	"net/http"
	"prog2-2025-dic-Heinen/dto"
	"prog2-2025-dic-Heinen/services"

	"github.com/gin-gonic/gin"
)

type ClimaHandler struct {
	service services.ClimaServiceInterface
}

func NewClimaHandler(ser services.ClimaServiceInterface) *ClimaHandler {
	return &ClimaHandler{
		service: ser,
	}
}

func (handler *ClimaHandler) CalcularMetricas(c *gin.Context) {
	var solicitud dto.TemperaturasRequest
	if err := c.ShouldBindJSON(&solicitud); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := handler.service.CalcularMetricas(solicitud)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (handler *ClimaHandler) CalcularProyeccion(c *gin.Context) {
	var solicitud dto.ProyeccionRequest
	if err := c.ShouldBindJSON(&solicitud); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := handler.service.CalcularProyeccion(solicitud)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (handler *ClimaHandler) ReporteOperaciones(c *gin.Context) {
	var solicitud dto.ReporteRequest
	if err := c.ShouldBindJSON(&solicitud); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := handler.service.ReporteOperaciones(solicitud)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (handler *ClimaHandler) RegistroEstaciones(c *gin.Context) {
	var solicitud dto.RegistroRequest
	if err := c.ShouldBindJSON(&solicitud); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := handler.service.RegistroEstaciones(solicitud)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
