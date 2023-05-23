package Models

import "time"

type AdminPassword struct {
	Password     string    `json:"password"`
	MainPassword string    `json:"mainPassword"`
	CreatedAt    time.Time `json:"created_at"`
}
