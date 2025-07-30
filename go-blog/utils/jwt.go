package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT wrapper
type JWTWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

// JWTClaim adds email as a claim to the token
type JWTClaim struct {
	Email string
	jwt.RegisteredClaims
}

func (j *JWTWrapper) GenerateToken(email string) (string, error) {
	/*
		generate a new JWT token for a given email
		- set the claims of the token
		- create the token with HS256 signing method
		- sign the token with secret key
		- return signed token
	*/

	// set the claims of the token
	claims := &JWTClaim{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(j.ExpirationHours))),
			Issuer:    j.Issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// create the token with hs256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign the token with secret key
	signedToken, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j *JWTWrapper) ValidateToken(signedToken string) (claims *JWTClaim, err error) {
	/*
		validate jwt token
		- parse the token with claims
		- check for parsing errors
		- validate token claims
		- check if token is expired
	*/

	// parse the token with claims
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)

	// check for parsing errors
	if err != nil {
		return
	}

	// validate token claims
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return nil, err
	}

	// check if token is expired
	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		return nil, err
	}

	return claims, nil
}

func NewJWTWrapper() *JWTWrapper {
	// JWTWrapper with configuration from env variables
	return &JWTWrapper{
		SecretKey: os.Getenv("JWT_SECRET"),
		Issuer: os.Getenv("JWT_ISSUER"),
		ExpirationHours: 24,
	}
}
