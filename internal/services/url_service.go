package services

import (
	"crypto/sha1"
	"encoding/base64"
	"strings"
	"time"
	"url-shortener/internal/models"
	"url-shortener/internal/repository"
)

func GenerateShortCode(LongURL string) string {
	hash := sha1.New()
	hash.Write([]byte(LongURL + time.Now().String()))
	sum := hash.Sum(nil)

	// Convert hash into letters/numbers (base64)
	code := base64.URLEncoding.EncodeToString(sum)

	// Take only first 6 characters (e.g., "abc123")
	return strings.TrimRight(code[:6], "=")
}

func CreateShortURL( longUrl string) (*models.URL, error) {
	shortCode := GenerateShortCode(longUrl)

	url := &models.URL{
		ID: shortCode,
		LongURL: longUrl,
		ShortURL: "http://localhost:8080/Bitly-lite/" + shortCode,
		CreatedAt: time.Now(),
		Clicks: 0,
	}

	// save in repository
	err := repository.SaveURL(url)
	if err != nil {
		return  nil, err
	}

	return url, nil
}

// GetLongURL fetches the original URL and increments clicks
func GetLongURL(code string) (string, error) {
	url, err := repository.FindURL(code)
	if err != nil {
		return "", err
	}

	// Update analytics (click counter)
	repository.IncrementClick(url.ID)

	return url.LongURL, nil
}