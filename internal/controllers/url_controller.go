package controllers

import (
	"net/http"
	"url-shortener/internal/services"

	"github.com/gin-gonic/gin"
)

func ShortenURL(c *gin.Context) {
	var body struct {
		LongURL string `json:"long_url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	url, err := services.CreateShortURL(body.LongURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to shorten URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"short_url": url.ShortURL,
		"clicks":    url.Clicks,
	})
}

func Redirect(c *gin.Context) {
	code := c.Param("code")

	longUrl, err := services.GetLongURL(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, longUrl)
}