package main

import (
	"url-shortener/internal/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/shorten", controllers.ShortenURL)
	r.GET("/:code", controllers.Redirect)

	r.Run(":8080")
}