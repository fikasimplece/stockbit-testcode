package models

import (
	"time"
)

type (
	// Movie
	Movie struct {
		ImdbID    string    `json:"imdbid"`
		Path      string    `json:"path`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
