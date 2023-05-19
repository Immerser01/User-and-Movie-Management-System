package LoginHandler

import (
	"github.com/Immerser01/InternAssignment/tree/main/Storage/Database"
	"github.com/gin-gonic/gin"
)

func LoginHandler(storage *Database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract username and password from the request body
		var user struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		// Perform authentication using the storage instance
		// ...

		c.JSON(200, gin.H{"message": "User logged in successfully"})
	}
}
