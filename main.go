package main

import (
	"url-shortener/internal/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/shorten", controllers.ShortenURL)
	r.GET("/Bitly-lite/:code", controllers.Redirect)
	r.GET("/analytics/:code", controllers.GetAnalytics)
	r.GET("/analytics/:code/events", controllers.GetEvents)


	r.Run(":8080")
}