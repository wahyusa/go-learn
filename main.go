package main

import (
	"go-learn/schema"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Average global var
var dbcon *gorm.DB

func main() {
	// Setup DB Connection
	dbcon = ConnectionStart()

	// generate tables
	dbcon.AutoMigrate(&schema.User{}, &schema.Photo{}, &schema.Comment{}, &schema.SocialMedia{})

}

func ConnectionStart() *gorm.DB {
	dsn := "host=localhost user=postgres password=password dbname=my_gram port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
