package LoginCredentials

import (
	"github.com/Immerser01/InternAssignment/tree/main/Verification/VerifyCredentials"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginCredentials(vdb *VerifyCredentials.Database) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		email := ginContext.PostForm("email")
		password := ginContext.PostForm("password")

		err := vdb.VerifyCredentials(email, password)
		if err != nil {
			ginContext.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ginContext.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	}
}
