package main

import (
	"database/sql"
	"github.com/Immerser01/InternAssignment/tree/main/Admin"
	CredentialHandler "github.com/Immerser01/InternAssignment/tree/main/Handler/CredentialHandler"
	_ "github.com/Immerser01/InternAssignment/tree/main/Handler/CredentialHandler"
	"github.com/Immerser01/InternAssignment/tree/main/Handler/Moviehandler"
	"github.com/Immerser01/InternAssignment/tree/main/Handler/UserHandler"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	// Database connection
	db, err := sql.Open("postgres", "postgres://sylvian-knight:Root2is1!414@localhost/InternAssignment?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	// Create user table if it doesn't exist
	_, err = db.Exec(`
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

	// Initialize Gin router
	r := gin.Default()

	// Create instances of the handlers
	userHandler := &UserHandler.UserHandler{
		DB: db,
	}
	movieHandler := &Moviehandler.MovieHandler{
		DB: db,
	}
	credentialHandler := &CredentialHandler.CredentialHandler{
		DB: db,
	}
	adminHandler := &Admin.AdminHandler{
		DB: db,
	}

	r.POST("/users", userHandler.CreateUser)
	r.POST("/credentials", credentialHandler.UpdateCredentials)
	r.GET("/users/:accessPassword", userHandler.ListUsers)
	r.GET("/AdminCredentialsPage", credentialHandler.ListCredentials)
	r.POST("/movies", movieHandler.AddMovie)
	r.DELETE("/movies/:id", movieHandler.DeleteMovie)
	r.GET("/movies/:id/:password", movieHandler.ListMoviesByUser)
	r.POST("/admin", adminHandler.PasswordManager)
	r.DELETE("admin/:password/:mainPassword", adminHandler.DeletePassword)
	r.GET("admin/:mainPassword", adminHandler.ListPassword)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

/*

curl -X POST -H "Content-Type: application/json" \
-d '{"name": "linuxize", "email": "linuxize@example.com", "dob": "2002-07-08}' \
https://localhost:8080/users

*/
