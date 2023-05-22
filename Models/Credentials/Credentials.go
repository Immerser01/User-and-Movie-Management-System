package Credentials

import "time"

type Credential struct {
	ID        int       `json:"id"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
