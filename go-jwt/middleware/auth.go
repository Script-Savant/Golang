package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("secret")

func JWTAuthMiddleware() gin.HandlerFunc {
	/*
	- check if the authorixation header is present and valid
	- extract the jwt token from the header
	- verify the tokensignature and expiration
	- if valid allow the request to continue
	- if invalid or missing, return 401
	*/

	return func(c *gin.Context) {
		// step 1 - get the authorization header (Bearer <token>)
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorizatioin header missing"})
			c.Abort()
			return
		}

		// step 2 - extract token from header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format (expected 'Bearer <token>')"})
			c.Abort()
			return 
		}

		// Step 3: Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ensure the token method is HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		// Step 4: Handle invalid token
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Step 5: Token is valid — proceed to the protected route
		c.Next()
	}
}