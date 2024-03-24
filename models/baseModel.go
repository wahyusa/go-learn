package models

import "time"

type baseModel struct {
	ID        uint       `gorm:"primmaryKey" json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
