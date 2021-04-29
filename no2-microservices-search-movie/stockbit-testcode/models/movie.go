package models

import (
	"time"
)

type (
	// Movie
	Movie struct {
		ID        int       `json:"id"`
		ImdbID    string    `json:"imdbid"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Path      string    `json:"path`
	}
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Movie
}
