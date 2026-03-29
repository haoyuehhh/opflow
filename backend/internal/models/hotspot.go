package models

import "time"

type Hotspot struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	Source    string    `json:"source"`
	Score     int       `json:"score"`
	CreatedAt time.Time `json:"created_at"`
}

type Content struct {
	ID        string    `json:"id"`
	HotspotID string    `json:"hotspot_id"`
	Platform  string    `json:"platform"`
	Text      string    `json:"text"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}