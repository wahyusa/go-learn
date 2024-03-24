package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBCON *gorm.DB

func ConnectionStart() *gorm.DB {
	var err error
	dsn := os.Getenv("DB")
	DBCON, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DBCON
}
