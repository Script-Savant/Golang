// JWT token creation based on user id and role

package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// GenerateToken -> generates a JWT token with user id and role
/*
1. Create claims including expiration
2. Sign the token using the secret key from env
3. Return the signed token
*/
func GenerateToken(userID uint, role string) (string, error) {
	// 1. Define JWT claims
	claims := jwt.MapClaims {
		"user_id": userID,
		"role": role,
		"exp": time.Now().Add(time.Hour * 14).Unix(),
	}

	// 2. Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 3. sign and return
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}