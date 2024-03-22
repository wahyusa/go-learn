package main

import (
	"go-learn/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Average global var
var dbcon *gorm.DB

func main() {
	// Setup DB Connection
	dbcon = ConnectionStart()

	// DEV ONLY auto truncate
	dbcon.Migrator().DropTable(&model.User{}, &model.Photo{}, &model.Comment{}, &model.SocialMedia{})

	// generate tables
	dbcon.AutoMigrate(&model.User{}, &model.Photo{}, &model.Comment{}, &model.SocialMedia{})

	// Gin router
	r := gin.Default()

	// Route group for `ping`
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Setup Gin",
		})
	})

	// Route group for `User` model
	users := r.Group("/users")
	{
		users.POST("/register", handleUserRegister)
		// users.POST("/login", handleUserLogin)
		// users.PUT("/users", handleUserUpdate)
		// users.DELETE("/users", handleUserDelete)
	}

	// Run on PORT 8080 as required in documentation
	r.Run()

}

func ConnectionStart() *gorm.DB {
	dsn := "host=localhost user=postgres password=password dbname=my_gram port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

// User controller

func handleUserRegister(c *gin.Context) {
	// User struct request and response format
	var newUserRequest model.User

	// Parse JSON request
	if err := c.ShouldBindJSON(&newUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the order to the database
	dbcon.Create(&newUserRequest)

	// Set response
	newUserResponse := model.UserRegisterResponse{
		ID:              newUserRequest.ID,
		Email:           newUserRequest.Email,
		Username:        newUserRequest.Username,
		Age:             newUserRequest.Age,
		ProfileImageURL: newUserRequest.ProfileImageURL,
	}

	c.JSON(http.StatusCreated, newUserResponse)
}
