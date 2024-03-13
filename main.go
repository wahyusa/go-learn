package main

import (
	"go-learn/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	dsn := "host=localhost user=postgres password=password dbname=basic_rest port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func handleGetOrders(c *gin.Context) {
	var ordersWithItems []model.Order
	dbcon.Preload(clause.Associations).Find(&ordersWithItems)
	c.JSON(http.StatusOK, ordersWithItems)
}

func handleCreateOrder(c *gin.Context) {
	var requestData model.Order

	// Parse JSON request
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the order to the database
	dbcon.Create(&requestData)

	c.JSON(http.StatusCreated, requestData)
}

func handleUpdateOrder(c *gin.Context) {
	orderID := c.Param("orderId")

	var updatedOrder model.Order

	// Parse JSON request directly into the Order struct
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the existing order in the database
	var existingOrder model.Order
	result := dbcon.Preload(clause.Associations).First(&existingOrder, orderID)

	// Check if the order exists
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Delete existing associated items
	dbcon.Where("order_id = ?", existingOrder.ID).Delete(&model.Item{})

	// Update the existing order with the new data
	existingOrder.OrderedAt = updatedOrder.OrderedAt
	existingOrder.CustomerName = updatedOrder.CustomerName
	existingOrder.Items = updatedOrder.Items

	// Save the updated order to the database
	dbcon.Save(&existingOrder)

	c.JSON(http.StatusOK, existingOrder)
}

func handleDeleteOrder(c *gin.Context) {
	orderIDStr := c.Param("orderId")

	// Convert order ID to uint
	orderID, err := strconv.ParseUint(orderIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	// Find the existing order in the database
	var existingOrder model.Order
	result := dbcon.Preload(clause.Associations).First(&existingOrder, uint(orderID))

	// Check if the order exists
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// // Hapus juga data "items" yang berelasi dgn order sampe ke akarnya
	for _, item := range existingOrder.Items {
		dbcon.Delete(&item)
	}

	// Delete the order from the database
	dbcon.Delete(&existingOrder)

	c.JSON(http.StatusOK, gin.H{"message": "Success delete"})
}
