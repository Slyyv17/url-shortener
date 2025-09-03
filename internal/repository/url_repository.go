package repository

import (
	"errors"
	"url-shortener/internal/models"
)

var urlStore = make(map[string]*models.URL)

// SaveURL stores a new short URL
func SaveURL(url *models.URL) error {
	urlStore[url.ID] = url
	return nil
}

// FindURL gets a URL by its short code
func FindURL(code string) (*models.URL, error) {
	if url, exists := urlStore[code]; exists {
		return url, nil
	}
	return nil, errors.New("URL not found")
}

// IncrementClick increases the click count
func IncrementClick(code string) {
	if url, exists := urlStore[code]; exists {
		url.Clicks++
	}
}