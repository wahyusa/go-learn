package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              int64  `gorm:"primaryKey"`
	Username        string `gorm:"size:50;not null"`
	Email           string `gorm:"size:150;not null"`
	Password        string `gorm:"type:text;not null"`
	Age             int    `gorm:"not null"`
	ProfileImageURL string `gorm:"type:text"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Photos          []Photo       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Comments        []Comment     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	SocialMedias    []SocialMedia `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type Photo struct {
	gorm.Model
	Title    string
	Caption  string
	PhotoURL string
	UserID   uint
	Comment  []Comment
}

type Comment struct {
	gorm.Model
	UserID  uint
	PhotoID uint
	Message string
}

type SocialMedia struct {
	gorm.Model
	Name           string
	SocialMediaURL string
	UserID         uint
}
