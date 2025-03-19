package models

import (
	"time"
)

// swagger:model
type Song struct {
	ID          uint      `json:"id" example:"1"`
	CreatedAt   time.Time `json:"created_at" example:"2025-03-19T00:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2025-03-19T00:00:00Z"`
	DeletedAt   time.Time `json:"deleted_at,omitempty" example:"null"`
	Group       string    `json:"group" gorm:"not null" example:"The Beatles"`
	Song        string    `json:"song" gorm:"not null" example:"Hey Jude"`
	ReleaseDate string    `json:"release_date" example:"1968-08-26"`
	Text        string    `json:"text" example:"Hey Jude, don't make it bad..."`
	Link        string    `json:"link" example:"https://example.com/heyjude"`
}
