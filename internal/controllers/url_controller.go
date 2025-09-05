package controllers

import (
	"net/http"
	"time"
	"url-shortener/internal/models"
	"url-shortener/internal/repository"
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

	// collect Analytics
	event := models.ClickEvent{
		URLID: code,
		Referrer: c.Request.Referer(),
		IP: c.ClientIP(),
		UserAgent: c.Request.UserAgent(),
		Country: c.GetHeader("X-Country"),
		Timestamp:  time.Now(),
	}

	// save event
	repository.SaveEvent(event)

	c.Redirect(http.StatusMovedPermanently, longUrl)
}

func GetAnalytics(c *gin.Context) {
	code := c.Param("code")

	clicks, err := services.GetAnalytics(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"shortCode": code,
		"clicks":    clicks,
	})
}

func GetEvents(c *gin.Context) {
	code := c.Param("code")

	events := repository.GetEvents(code)

	c.JSON(http.StatusOK, gin.H{
		"shortCode": code,
		"events":    events,
	})
}