package controllers

import (
	"net/http"
	"time"
	"url-shortener/internal/models"
	"url-shortener/internal/repository"
	"url-shortener/internal/services"

	"github.com/gin-gonic/gin"
)

// ShortenURL godoc
// @Summary      Shorten a URL
// @Description  Takes a long URL and returns a shortened version
// @Tags         shorten
// @Accept       json
// @Produce      json
// @Param        request body models.ShortenRequest true "Shorten request"
// @Success      200 {object} models.ShortenResponse
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /shorten [post]
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

// Redirect godoc
// @Summary      Redirect shortened URL
// @Description  Redirects to the original long URL using the short code
// @Tags         redirect
// @Param        code path string true "Short code"
// @Success      301 {string} string "Redirected to original URL"
// @Failure      404 {object} map[string]string
// @Router       /{code} [get]
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

// GetAnalytics godoc
// @Summary Get click count for a short URL
// @Description Returns the number of times a short URL has been accessed
// @Tags analytics
// @Produce json
// @Param code path string true "Short code"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]string
// @Router /analytics/{code} [get]
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

// GetEvents godoc
// @Summary Get detailed analytics events
// @Description Returns detailed analytics including referrer, IP, user-agent, and timestamp
// @Tags analytics
// @Produce json
// @Param code path string true "Short code"
// @Success 200 {object} map[string]interface{}
// @Router /analytics/{code}/events [get]
func GetEvents(c *gin.Context) {
	code := c.Param("code")

	events := repository.GetEvents(code)

	c.JSON(http.StatusOK, gin.H{
		"shortCode": code,
		"events":    events,
	})
}