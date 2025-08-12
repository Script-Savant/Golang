/*
 Purpose: JWT authentication middleware
 Pseudo-code:
   - parse Authorization header Bearer token
   - validate token signature and expiration
   - attach user ID to gin.Context
*/
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims is the custom JWT claims we use.
type JWTClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// JWTAuthMiddleware returns a gin middleware that validates JWTs.
func JWTAuthMiddleware(secret string) gin.HandlerFunc {
	// Steps:
	// 1. Read Authorization header and extract bearer token
	// 2. Parse and validate token with jwt.ParseWithClaims
	// 3. If valid, set user_id in context; otherwise abort 401
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if h == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			return
		}
		// Expect "Bearer <token>"
		var tokenStr string
		_, _ = fmt.Sscanf(h, "Bearer %s", &tokenStr)
		if tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "bearer token required"})
			return
		}

		// Parse token
		token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin