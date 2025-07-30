package middleware

import (
	"go-blog/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Step 1: Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Step 2: Check if the header is in the format "Bearer <token>"
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		// Step 3: Validate the token
		token := headerParts[1]
		jwtWrapper := utils.NewJWTWrapper()
		claims, err := jwtWrapper.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Step 4: Set the email in the context for use in subsequent handlers
		c.Set("email", claims.Email)
		c.Next()
	}
}