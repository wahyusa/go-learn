package controllers

import (
	"go-learn/config"
	"go-learn/helpers"
	"go-learn/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

func LoginAttempt(c *gin.Context) {
	var jsonRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.BindJSON(&jsonRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	// cari email dulu kalo gk ada gak usah lanjut cek pw
	config.DBCON.First(&user, "email = ?", jsonRequest.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect user / pw"})
		return
	}

	// cek password hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(jsonRequest.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect user / pw"})
		return
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenResult, err := token.SignedString([]byte(os.Getenv("NOT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fail to generate token"})
		return
	}

	c.JSON(200, gin.H{"token": tokenResult})
}
