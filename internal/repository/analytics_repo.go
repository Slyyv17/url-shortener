// repository/analytics_repository.go
package repository

import (
	"url-shortener/internal/models"
)

var events = []models.ClickEvent{}

// SaveEvent saves an analytics event
func SaveEvent(event models.ClickEvent) {
	events = append(events, event)
}

// GetEvents returns all analytics for a given short code
func GetEvents(code string) []models.ClickEvent {
	var filtered []models.ClickEvent
	for _, e := range events {
		if e.URLID == code {
			filtered = append(filtered, e)
		}
	}
	return filtered
}
