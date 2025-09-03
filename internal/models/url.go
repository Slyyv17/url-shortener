// Model for a Shortened URL
package models

import "time"

type URL struct {
	ID        string `json:"id"`
	ShortURL  string `json:"short_url"`
	LongURL   string `json:"long_url"`
	CreatedAt time.Time `json:"created_at"`
	Clicks    int       `json:"clicks"`
}