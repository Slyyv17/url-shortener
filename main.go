package main

import (
	"os"
	_ "url-shortener/docs"
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
	r := gin.Default()
	r.POST("/shorten", controllers.ShortenURL)
	r.GET("/Bitly-lite/:code", controllers.Redirect)
	r.GET("/analytics/:code", controllers.GetAnalytics)
	r.GET("/analytics/:code/events", controllers.GetEvents)

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default for local dev
	}

	r.Run(":" + port)
}