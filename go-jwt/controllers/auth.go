package controllers

import (
	"go-jwt/config"
	"go-jwt/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context){
	/*
	receive json payload from the request
	validate input
	hash the password
	save the new user to the db
	*/

	var user models.User

	// bind json payload to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		// return error if json is invalid
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// hash the user's pass using bcrypt
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashin password"})
		return
	}

	// replace plain-text password with hashed vesrion
	user.Password = string(hashedPass)

	// save the user to the database
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user" : gin.H{
			"id": user.ID,
			"name": user.Name,
			"email": user.Email,
		},
	})
}

var jwtSecret = []byte("secret")

func LoginUser(c *gin.Context){
 /*
 - parse the incoming json request body to extract user's email and password
 - check if the user with the provided email exists
 - compare the provided password with te hashed password stored in the database
 - if authentication is successful, generate a JWT token containing the user's id and expiration timestamp
 - sign the jwt using a secret to ensure it's tamper proof
 - return success response containing the signed jwt, which the client can use to access
 */

 var input models.User // temporarily hold the incoming login data
 var user models.User // hold the actual user fetched from the database

 // step 1 - parse json input
 if err := c.ShouldBindJSON(&input); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
 }

 // step 2 - look up the user by email
 if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
	return
 }

 // step 3 - compare password using bcrypt
 if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Inavalid email or password"})
	return
 }

 // step 4 - create jwt with user_id and expiration
 token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"user_id": user.ID,
	"exp": time.Now().Add(time.Hour * 72).Unix(), // token valid for 72 hours
 })

 // step 5 - sign the token with secret key
 tokenString, err := token.SignedString(jwtSecret)
 if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
	return
 }

 // step 6 - return the token in the response
 c.JSON(http.StatusOK, gin.H{
	"message": "Login successful",
	"token": tokenString,
 })
 
}