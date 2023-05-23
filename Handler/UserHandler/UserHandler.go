package UserHandler

import (
	"database/sql"
	"github.com/Immerser01/InternAssignment/tree/main/Models/User"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

type UserHandler struct {
	DB *sql.DB
}

// CreateUser creates a new user
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user User.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Create user error"})
		return
	}

	if user.Email == "" || user.Name == "" {
		c.JSON(500, gin.H{
			"ERROR": "None of the Credentials fields can be empty!"})
		return
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(user.Email) {
		c.JSON(501, gin.H{"error": "Invalid email format"})
		return
	}

	DOBRegex := regexp.MustCompile(`^\d{2}-\d{2}-\d{4}$`)
	if !DOBRegex.MatchString(user.DOB) {
		c.JSON(502, gin.H{"error": "Invalid date of birth format"})
		return
	}

	// Insert user into the database
	query := `
		INSERT INTO UserData (email, name, dob)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`
	err := h.DB.QueryRow(query, user.Email, user.Name, user.DOB).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Create user error 2"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// ListUsers lists all registered users
func (h *UserHandler) ListUsers(c *gin.Context) {
	accessPassword := c.Param("password")

	c.JSON(http.StatusOK, gin.H{
		"accessPassword": accessPassword,
	})
	query := `
		SELECT id, email, name, dob, created_at FROM UserData;
	`
	rows, err := h.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in List Users function"})
		return
	}
	defer rows.Close()

	var users []User.User
	for rows.Next() {
		var user User.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.DOB, &user.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "LUE2"})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}
