package models

type User struct {
	Model
	Email           string `json:"email" gorm:"not null;unique" binding:"required"`
	Username        string `json:"username" gorm:"not null;unique" binding:"required"`
	Age             int    `json:"age" gorm:"not null"  binding:"required"`
	Password        string `json:"password" gorm:"not null"  binding:"required"`
	ProfileImageURL string `json:"profile_image_url" gorm:"default:null"`
	// Photos          []Photo       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	// Comments        []Comment     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	// SocialMedias    []SocialMedia `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
