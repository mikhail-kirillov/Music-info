package models

import (
	"gorm.io/gorm"
)

type SongTable struct {
	gorm.Model
	Group       string `json:"group" gorm:"not null"`
	Song        string `json:"song" gorm:"not null"`
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

func (SongTable) TableName() string {
	return "songs"
}

// swagger:model
type SongResponse struct {
	ID          uint   `json:"id" example:"1"`
	Group       string `json:"group" example:"The Beatles"`
	Song        string `json:"song" example:"Hey Jude"`
	ReleaseDate string `json:"release_date" example:"1968-08-26"`
	Text        string `json:"text" example:"Hey Jude, don't make it bad..."`
	Link        string `json:"link" example:"https://example.com/heyjude"`
}
