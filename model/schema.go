package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string
	Email       string
	Password    string
	Age         uint
	Photo       []Photo
	Comment     []Comment
	SocialMedia []SocialMedia
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
