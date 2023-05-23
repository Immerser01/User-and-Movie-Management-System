package Models

import (
	"database/sql"
	"log"
)

func CreateTables(db *sql.DB) {

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS UserData (
			id SERIAL PRIMARY KEY,
			email VARCHAR(255) NOT NULL,
			name VARCHAR(255) NOT NULL,
			dob VARCHAR(255) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW()
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create movie table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS MovieData (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			title VARCHAR(255) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW()
			
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Credential (
			id SERIAL PRIMARY KEY,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW()
			
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Admin (
		    
		    mainPassword VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW()
			
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
}
