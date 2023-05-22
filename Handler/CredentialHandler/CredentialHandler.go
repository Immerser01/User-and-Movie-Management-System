package CredentialHandler

import (
	"database/sql"
	"github.com/Immerser01/InternAssignment/tree/main/Models/Credentials"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CredentialHandler struct {
	DB *sql.DB
}

func (h *CredentialHandler) UpdateCredentials(c *gin.Context) {
	var credential Credentials.Credential
	if err := c.ShouldBindJSON(&credential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Create Credential connecting error"})
		return
	}

	if credential.Password == "" {
		c.JSON(500, gin.H{
			"ERROR": "None of the Credentials fields can be empty!"})
		return
	}

	// Insert user into the database
	query := `
		INSERT INTO Credential (id, password)
		VALUES ($1, $2)
		RETURNING created_at
	`
	err := h.DB.QueryRow(query, credential.ID, credential.Password).Scan(&credential.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Duplicate ID. Please don't do this."})
		return
	}

	c.JSON(http.StatusCreated, credential)
}

// Only for Administrator
func (h *CredentialHandler) ListCredentials(c *gin.Context) {
	rows, err := h.DB.Query("SELECT id, password FROM Credential")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Credential listing error"})
		return
	}
	defer rows.Close()

	var cred []Credentials.Credential
	for rows.Next() {
		var credential Credentials.Credential
		if err := rows.Scan(&credential.ID, &credential.Password); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Credential accessing error"})
			return
		}
		cred = append(cred, credential)
	}

	c.JSON(http.StatusOK, cred)
}
