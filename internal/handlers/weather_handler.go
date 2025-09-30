package handlers

import (
	"net/http"

	"cloud-run-cep/internal/services"

	"github.com/gin-gonic/gin"
)

type WeatherResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

type WeatherHandler struct {
	cepService     services.CEPService
	weatherService services.WeatherService
}

func NewWeatherHandler(cepService services.CEPService, weatherService services.WeatherService) *WeatherHandler {
	return &WeatherHandler{
		cepService:     cepService,
		weatherService: weatherService,
	}
}

func (h *WeatherHandler) GetWeather(c *gin.Context) {
	zipcode := c.Param("zipcode")

	location, err := h.cepService.GetLocationByCEP(zipcode)
	if err != nil {
		switch err.Error() {
		case "invalid zipcode":
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid zipcode"})
			return
		case "can not find zipcode":
			c.JSON(http.StatusNotFound, gin.H{"error": "can not find zipcode"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
	}

	tempC, tempF, tempK, err := h.weatherService.GetWeatherByLocation(location)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "can not find zipcode"})
		return
	}

	response := WeatherResponse{
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	}

	c.JSON(http.StatusOK, response)
}
