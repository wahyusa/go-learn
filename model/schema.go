package model

import (
	"time"
)

type Order struct {
	ID           uint      `json:"id"`
	OrderedAt    time.Time `json:"orderedAt"`
	CustomerName string    `json:"customerName"`
	Items        []Item    `gorm:"foreignKey:OrderID" json:"items"`
}

type Item struct {
	ID          uint   `json:"-"`
	Code        string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"-"`
}
