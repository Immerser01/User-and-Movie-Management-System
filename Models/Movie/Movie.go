package Movie

import (
	"time"
)

// Movie struct to represent a movie
type Movie struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}
