package Moviehandler

import (
	"database/sql"
	"github.com/Immerser01/User-and-Movie-Management-System/tree/main/Models/Credentials"
	"github.com/Immerser01/User-and-Movie-Management-System/tree/main/Models/Movie"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MovieHandler struct {
	DB *sql.DB
}

//type Movie

// AddMovie adds a new movie
func (h *MovieHandler) AddMovie(c *gin.Context) {
	var movie Movie.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "AME1"})
		return
	}

	// Insert movie into the database
	query := `
		INSERT INTO MovieData (user_id, title)
		VALUES ($1, $2)
		RETURNING id, created_at
	`
	err := h.DB.QueryRow(query, movie.UserID, movie.Title).Scan(&movie.ID, &movie.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AME2"})
		return
	}

	c.JSON(http.StatusCreated, movie)
}

// DeleteMovie deletes a movie by ID
func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Delete movie from the database
	result, err := h.DB.Exec("DELETE FROM MovieData WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DME1"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DME2"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}

// ListMoviesByUser lists movies watched by a user
func (h *MovieHandler) ListMoviesByUser(c *gin.Context) {

	var credential Credentials.Credential
	stringUserID := c.Param("id")
	intUserID, err := strconv.Atoi(stringUserID)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	credential.ID = intUserID
	credential.Password = c.Param("password")

	c.JSON(http.StatusOK, gin.H{
		"id":       credential.ID,
		"password": credential.Password,
	})

	query := `
		SELECT id, user_id, title, created_at FROM MovieData WHERE user_id = $3 AND EXISTS (SELECT id, password FROM Credential WHERE id = $1 AND password = $2);
	`

	rows, err1 := h.DB.Query(query, credential.ID, credential.Password, credential.ID)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var movies []Movie.Movie
	for rows.Next() {
		var movie Movie.Movie
		if err := rows.Scan(
			&movie.ID,
			&movie.UserID,
			&movie.Title,
			&movie.CreatedAt,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "LME3"})
			return
		}
		movies = append(movies, movie)

	}
	if len(movies) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "You are an imposter who entered wrong password. Or you dont even exist. Or maybe you didnt enter any movies. Either way, you stupid."})
		return
	}
	c.JSON(http.StatusOK, movies)
}
