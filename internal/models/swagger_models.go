package models

// ShortenRequest is the body for POST /shorten
type ShortenRequest struct {
	LongURL string `json:"long_url" example:"https://www.google.com/" binding:"required"`
}

// ShortenResponse is returned after shortening
type ShortenResponse struct {
	ShortURL string `json:"short_url" example:"http://localhost:8080/abc123"`
	Clicks   int    `json:"clicks" example:"0"`
}