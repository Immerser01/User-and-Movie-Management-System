package Admin

import (
	"database/sql"
	"github.com/Immerser01/InternAssignment/tree/main/Models"
	"github.com/Immerser01/InternAssignment/tree/main/Models/Credentials"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CentralPassword = "ThisIsMainPassword"
)

type AdminHandler struct {
	DB *sql.DB
}

func (h *AdminHandler) PasswordManager(c *gin.Context) {

	var ps Models.AdminPassword
	if err := c.ShouldBindJSON(&ps); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "admin server error"})
		return
	}

	if ps.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"eror": "Password can't be blank"})
		return
	}
	if ps.MainPassword == CentralPassword {
		query := `
			INSERT INTO Admin (mainPassword, password) 
			VALUES ($1, $2)
			RETURNING created_at
		`

		err := h.DB.QueryRow(query, ps.MainPassword, ps.Password).Scan(&ps.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Password Creating Error"})
			return
		}

		c.JSON(http.StatusCreated, ps)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The Main Password you entered in incorrect"})
		return
	}
}

func (h *AdminHandler) ListPassword(c *gin.Context) {
	mainPassword := c.Param("mainPassword")

	c.JSON(http.StatusOK, gin.H{
		"mainPassword": mainPassword,
	})

	if mainPassword == CentralPassword {
		rows, err := h.DB.Query("SELECT mainPassword, password, created_at FROM Admin")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't execute query"})
			return
		}
		defer rows.Close()

		var pass []Models.AdminPassword
		for rows.Next() {
			var password Models.AdminPassword
			if err := rows.Scan(
				&password.Password,
				&password.MainPassword,
				&password.CreatedAt,
			); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Scanning error"})
				return
			}
			pass = append(pass, password)

		}
		if len(pass) == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No passwords here"})
			return
		}
		c.JSON(http.StatusOK, pass)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "main password wrong"})
		return
	}
}

func (h *AdminHandler) DeletePassword(c *gin.Context) {
	password := c.Param("password")
	mainPassword := c.Param("mainPassword")

	c.JSON(http.StatusOK, gin.H{
		"password":     password,
		"mainPassword": mainPassword,
	})

	if mainPassword == CentralPassword {
		result, err := h.DB.Exec("DELETE FROM Admin WHERE password = $1", password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't execute query"})
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't execute Row query"})
			return
		}

		if rowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No such password found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Done masterfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "main password wrong"})
		return
	}
}

// Only for Administrator
func (h *AdminHandler) ListCredentials(c *gin.Context) {
	mainPassword := c.Param("mainPassword")

	c.JSON(http.StatusOK, gin.H{
		"mainPassword": mainPassword,
	})

	if mainPassword == CentralPassword {
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
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Main Password"})
	}
}
