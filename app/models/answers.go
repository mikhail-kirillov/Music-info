package models

// swagger:model
type GetSongsResponse struct {
	Data       []SongResponse `json:"data"`
	Page       int            `json:"page" example:"1"`
	Limit      int            `json:"limit" example:"10"`
	TotalPages int            `json:"totalPages" example:"5"`
	TotalItems int64          `json:"totalItems" example:"50"`
}

// swagger:model
type ErrorResponse struct {
	Error string `json:"error" example:"Invalid request"`
}

// swagger:model
type GetSongLyricsResponse struct {
	Verses      []string `json:"verses"`
	Page        int      `json:"page" example:"1"`
	Limit       int      `json:"limit" example:"10"`
	TotalVerses int      `json:"totalVerses" example:"50"`
}

type AddSongRequest struct {
	Group string `json:"group" example:"The Beatles" binding:"required"`
	Song  string `json:"song" example:"Hey Jude" binding:"required"`
}

// swagger:model
type UpdateSongRequest struct {
	Group       *string `json:"group,omitempty" example:"The Beatles"`
	Song        *string `json:"song,omitempty" example:"Hey Jude"`
	ReleaseDate *string `json:"release_date,omitempty" example:"1968-08-26"`
	Text        *string `json:"text,omitempty" example:"Hey Jude, don't make it bad..."`
	Link        *string `json:"link,omitempty" example:"https://example.com/heyjude"`
}

// swagger:model
type DeleteSongResponse struct {
	Message string `json:"message" example:"Song deleted"`
}
