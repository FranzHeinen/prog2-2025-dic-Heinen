package services

import (
	"errors"
	"math"
	"prog2-2025-dic-Heinen/dto"
	"time"
)

type ClimaServiceInterface interface {
	CalcularMetricas(datos dto.TemperaturasRequest) (dto.TemperaturasResponse, error)
	CalcularProyeccion(datos dto.ProyeccionRequest) ([]dto.ProyeccionResponse, error)
	ReporteOperaciones(datos dto.ReporteRequest) ([]dto.ReporteResponse, error)
	RegistroEstaciones(datos dto.RegistroRequest) (dto.RegistroResponse, error)
}

type ClimaService struct{}

func NewClimaService() *ClimaService {
	return &ClimaService{}
}

func (service *ClimaService) CalcularMetricas(datos dto.TemperaturasRequest) (dto.TemperaturasResponse, error) {

	var suma float64
	var diferencia float64
	var cantidad int
	var sumaDiferencias float64

	for i := 0; i < len(datos.Temperaturas); i++ {
		suma = datos.Temperaturas[i]
		cantidad++
	}

	promedio := suma / float64(cantidad)
	mediana := datos.Temperaturas[cantidad/2]

	for i := 0; i < len(datos.Temperaturas); i++ {
		diferencia = promedio - datos.Temperaturas[i]

		diferencia = diferencia * diferencia

		sumaDiferencias += diferencia
	}

	desvioEstandar := sumaDiferencias / float64(cantidad)

	desvioEstandar = math.Sqrt(desvioEstandar)

	return dto.TemperaturasResponse{
		Promedio:       promedio,
		Mediana:        mediana,
		DesvioEstandar: desvioEstandar,
		Unidad:         "Celsius",
	}, nil
}

func (service *ClimaService) CalcularProyeccion(datos dto.ProyeccionRequest) ([]dto.ProyeccionResponse, error) {

	var resultado []dto.ProyeccionResponse
	var valorAnterior float64
	var temperaturaActual float64
	var anio int
	var fecha = datos.Fecha

	for i := 0; i < datos.Anios; i++ {

		valorAnterior = datos.TemperaturaActual

		anio = i + 1
		calculo := valorAnterior * datos.TasaAnual / 100

		temperaturaActual = valorAnterior + calculo

		resultado = append(resultado, dto.ProyeccionResponse{Anio: anio, Temperatura: temperaturaActual, Fecha: fecha})
	}
	return resultado, nil
}

//func (service *ClimaService) ReporteOperaciones(datos dto.ReporteRequest, parametros dto.ProyeccionRequest) ([]dto.ReporteResponse, error) {

//var fechaDesde = datos.FechaDesde
//var fechaHasta = datos.FechaHasta
//}

func (service *ClimaService) RegistroEstaciones(datos dto.RegistroRequest) (dto.RegistroResponse, error) {
	if datos.Nombre != "" && datos.Latitud >= -90 && datos.Latitud <= 90 && datos.Longitud >= 180 && datos.Longitud <= 180 && datos.Altitud > 0 {
		return dto.RegistroResponse{}, errors.New("Error: datos inválidos, Detalles: latitud fuera de rango, Nombre: nombre es obligatorio")
	}
	var id int

	id = id + 1

	return dto.RegistroResponse{Id: id, Nombre: datos.Nombre, Latitud: datos.Latitud, Longitud: datos.Longitud, Altitud: datos.Altitud, FechaCreacion: time.Now()}, nil

}
