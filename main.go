package main

import (
	"go-learn/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Average global var
var dbcon *gorm.DB

func main() {
	// Setup DB Connection
	dbcon = ConnectionStart()

	// generate tables
	dbcon.AutoMigrate(&model.User{}, &model.Photo{}, &model.Comment{}, &model.SocialMedia{})

}

func ConnectionStart() *gorm.DB {
	dsn := "host=localhost user=postgres password=password dbname=my_gram port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
