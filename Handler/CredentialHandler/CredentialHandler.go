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
		if err := rows.Scan(&credential.ID, &credential.Password, &credential.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Credential accessing error"})
			return
		}
		cred = append(cred, credential)
	}

	c.JSON(http.StatusOK, cred)
}

//query2 := `
//	INSERT INTO Credential (id, password)
//	VALUES ($1, $2)
//	RETURNING created_at
//`
//err = h.DB.QueryRow(query2, credential.ID, credential.Password).Scan(&user.CreatedAt)
//if err != nil {
//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Credential table error"})
//	return
//}
