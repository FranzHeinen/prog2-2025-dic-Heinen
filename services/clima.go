package services

import (
	"math"
	"prog2-2025-dic-Heinen/dto"
)

type ClimaServiceInterface interface {
	CalcularMetricas(dto.TemperaturasRequest) (dto.TemperaturasResponse, error)
	CalcularProyeccion(dto.ProyeccionRequest) ([]dto.ProyeccionResponse, error)
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
		Unidad:         "celsius",
	}, nil
}

func (service *ClimaService) CalcularProyeccion(datos dto.ProyeccionRequest) ([]dto.ProyeccionResponse, error) {

	var resultado []dto.ProyeccionResponse
	var valorAnterior float64
	var temperaturaActual float64
	var anio int

	for i := 0; i < datos.Anios; i++ {

		valorAnterior = datos.TemperaturaActual

		anio = i + 1
		calculo := valorAnterior * datos.TasaAnual / 100

		temperaturaActual = valorAnterior + calculo

		resultado = append(resultado, dto.ProyeccionResponse{Anio: anio, Temperatura: temperaturaActual})
	}
	return resultado, nil
}

func (service *ClimaService) ReporteOperaciones() {

}

func (service *ClimaService) RegistroEstaciones() {

}
