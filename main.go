package main

import (
	"log"
	"os"
	"url-shortener/docs"
	"url-shortener/internal/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Bitly-lite URL Shortener API
// @version 1.0
// @description A simple URL shortener with analytics built using Gin.
// @host localhost:8080
// @BasePath /
func main() {
	// Auto-detect environment and set Swagger host
	if os.Getenv("RENDER") != "" {
		// Running on Render
		docs.SwaggerInfo.Host = "byte-lite-be6i.onrender.com"
		docs.SwaggerInfo.Schemes = []string{"https"}
	} else {
		// Running locally
		docs.SwaggerInfo.Host = "localhost:8080"
		docs.SwaggerInfo.Schemes = []string{"http"}
	}
	
	r := gin.Default()
	r.POST("/shorten", controllers.ShortenURL)
	r.GET("/Bitly-lite/:code", controllers.Redirect)
	r.GET("/analytics/:code", controllers.GetAnalytics)
	r.GET("/analytics/:code/events", controllers.GetEvents)

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}