package main

import (
	"go-learn/model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Using gorm example usage
var dbcon *gorm.DB

func main() {
	dbcon = OpenConnection()

	// auto bikin table
	dbcon.AutoMigrate(&model.Order{}, &model.Item{})

	// setup gin
	router := gin.Default()

	// Order routes
	router.GET("/orders", handleGetOrders)
	router.POST("/orders", handleCreateOrder)
	router.PUT("/orders/:orderId", handleUpdateOrder)
	router.DELETE("/orders/:orderId", handleDeleteOrder)

	// custom port
	router.Run(":6969")
}

func OpenConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=password dbname=gormpzn port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
