/*
JWT authentication middleware for protected routes
1. Extract token from authorization header
2. validate token an dparse claims
3. Set user id in context fro downstream handlers
*/

package middleware

import (
	"ecommerce-api/internal/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Get and validate authorization header
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			unauthorized(c, "Authorization header missing")
			return
		}

		// 2. Extract token string
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 3. parse and validate token
		secret := []byte(config.LoadConfig().JWTSecret)
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return secret, nil
		})
		if err != nil || !token.Valid {
			unauthorized(c, "Invalid token")
			return
		}

		// 4. Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			unauthorized(c, "Invalid token claims")
			return
		}

		// 5. Set user ID in context
		if uid, ok := claims["user_id"].(float64); ok {
			c.Set("userID", uint(uid))
			c.Next()
		} else {
			unauthorized(c, "User ID not found in token")
		}
	}
}

func unauthorized(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": msg})
}
