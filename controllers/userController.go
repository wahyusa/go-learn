package controllers

import (
	"go-learn/config"
	"go-learn/helpers"
	"go-learn/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func StoreUser(c *gin.Context) {
	var jsonRequest struct {
		Email           string `json:"email" binding:"required"`
		Username        string `json:"username" binding:"required"`
		Age             int    `json:"age" binding:"required"`
		Password        string `json:"password" binding:"required"`
		ProfileImageURL string `json:"profile_image_url"`
	}

	if err := c.BindJSON(&jsonRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Password hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(jsonRequest.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
		return
	}

	// cek panjang char password
	if len(jsonRequest.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password minimal 6 chars"})
		return
	}

	// cek value age minimal 8
	if jsonRequest.Age < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Age minimal 8 years old"})
		return
	}

	// format email harus valid
	if !helpers.IsValidEmail(jsonRequest.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address"})
		return
	}

	// profile_image_URL harus valid
	if !helpers.IsValidURL(jsonRequest.ProfileImageURL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL format"})
		return
	}
	// username harus unique
	if !helpers.IsUniqueUsername(config.DBCON, jsonRequest.Username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	// email harus unique
	if !helpers.IsUniqueEmail(config.DBCON, jsonRequest.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	// Insert to DB
	user := models.User{
		Email:           jsonRequest.Email,
		Username:        jsonRequest.Username,
		Age:             int(jsonRequest.Age),
		Password:        string(hashedPassword),
		ProfileImageURL: jsonRequest.ProfileImageURL,
	}

	result := config.DBCON.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to insert"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User stored successfully"})
}
