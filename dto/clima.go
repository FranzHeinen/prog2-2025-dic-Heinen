package dto

import "time"

type TemperaturasRequest struct {
	Temperaturas []float64 `json:"temperaturas"`
}

type TemperaturasResponse struct {
	Promedio       float64 `json:"promedio"`
	Mediana        float64 `json:"mediana"`
	DesvioEstandar float64 `json:"desvio_estandar"`
	Unidad         string  `json:"unidad"`
}

type ProyeccionRequest struct {
	TemperaturaActual float64   `json:"temperatura_actual"`
	TasaAnual         float64   `json:"tasa_anual"`
	Anios             int       `json:"anios"`
	Fecha             time.Time `json:"fecha"`
}

type ProyeccionResponse struct {
	Anio        int       `json:"anio"`
	Temperatura float64   `json:"temperatura"`
	Fecha       time.Time `json:"fecha"`
}

type ReporteRequest struct {
	FechaDesde time.Time `json:"fecha_desde"`
	FechaHasta time.Time `json:"fecha_hasta"`
}

type ReporteResponse struct {
	Tipo      float64   `json:"tipo"`
	Fecha     time.Time `json:"fecha"`
	Resultado float64   `json:"resultado"`
}

type RegistroRequest struct {
	Nombre   string  `json:"nombre"`
	Latitud  float64 `json:"latitud"`
	Longitud float64 `json:"longitud"`
	Altitud  float64 `json:"altitud"`
}

type RegistroResponseFalla struct {
	Error    string `json:"error"`
	Detalles string `json:"detalles"`
}

type RegistroResponse struct {
	Id            string    `json:"id"`
	Nombre        string    `json:"nombre"`
	Latitud       float64   `json:"latitud"`
	Longitud      float64   `json:"longitud"`
	Altitud       float64   `json:"altitud"`
	FechaCreacion time.Time `json:"fecha_creacion"`
}
