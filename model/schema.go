package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              int
	Email           string `gorm:"size:150;not null"`
	Username        string `gorm:"size:50;not null"`
	Age             int    `gorm:"not null"`
	Password        string `gorm:"type:text;not null"`
	ProfileImageURL string `gorm:"type:text"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Photos          []Photo       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Comments        []Comment     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	SocialMedias    []SocialMedia `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type UserRegisterResponse struct {
	ID              int    `json:"id"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	Age             int    `json:"age"`
	ProfileImageURL string `json:"profile_image_url"`
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
