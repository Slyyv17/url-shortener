// model for analytics
package models

import "time"

type ClickEvent struct {
	URLID     string    `json:"url_id"`
	Referrer  string    `json:"referrer"`
	IP        string    `json:"ip"`
	UserAgent string    `json:"user-agent"`
	Country   string    `json:"country"`
	Timestamp time.Time `json:"timestamp"`
}