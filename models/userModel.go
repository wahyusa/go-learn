package models

type User struct {
	Model
	Email           string `json:"email" gorm:"not null;unique" binding:"required"`
	Username        string `json:"username" gorm:"not null;unique" binding:"required"`
	Age             uint   `json:"age" gorm:"not null"  binding:"required"`
	Password        string `json:"password" gorm:"not null"  binding:"required"`
	ProfileImageURL string `json:"profile_image_url"`
	// Photos          []Photo       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	// Comments        []Comment     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	// SocialMedias    []SocialMedia `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
