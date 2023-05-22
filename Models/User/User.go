package User

import "time"

// User struct to represent a registered user
type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	DOB       string    `json:"dob"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
