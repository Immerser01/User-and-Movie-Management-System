package Authorization

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware represents the authentication middleware
type AuthMiddleware struct {
}

// NewAuthMiddleware creates a new instance of AuthMiddleware
func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

// Authenticate is the middleware function for authentication
func (am *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Extract the token from the header
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Parse the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}

			//To be changed according to the actual secret key
			secretKey := []byte("mDYOgzDLO2zIXw1Ip2ewL20EfMKLm8F1h2XZ3JRpYZZemkGIf8SpbCx9bO22EwQ\n")
			return secretKey, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Verify the token
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Store the user ID in the request context for further use
		claims := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"].(float64) // Assuming the user ID is stored in the "user_id" claim as a float64
		c.Set("user_id", int(userID))

		// Proceed to the next middleware or handler
		c.Next()
	}
}
