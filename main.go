package main

import (
	"log"
	"os"

	"cloud-run-cep/internal/handlers"
	"cloud-run-cep/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	cepService := services.NewCEPService()
	weatherService := services.NewWeatherService()

	weatherHandler := handlers.NewWeatherHandler(cepService, weatherService)

	router := gin.Default()
	router.GET("/weather/:zipcode", weatherHandler.GetWeather)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor iniciado na porta %s", port)
	log.Fatal(router.Run(":" + port))
}
