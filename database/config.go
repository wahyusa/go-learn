package database

import (
	"go-learn/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbcon *gorm.DB

func ConnectionStart() *gorm.DB {
	dsn := "host=localhost user=postgres password=password dbname=my_gram port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// DEV ONLY auto truncate
	db.Migrator().DropTable(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})

	// generate tables
	db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})
	return db
}
