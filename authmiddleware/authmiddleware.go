package authmiddleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Fetch the JWT secret key from environment variable
var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from the Authorization header
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}
		// Remove "Bearer " from the token string
		tokenStr = tokenStr[len("Bearer "):]

		// Define claims structure
		claims := &jwt.RegisteredClaims{}

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil // Return the secret key for validation
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Store the user's ID or other relevant claim in the context
		c.Set("user", claims.Subject)

		// Proceed to the next middleware/handler
		c.Next()
	}
}
