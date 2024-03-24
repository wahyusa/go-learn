package models

type User struct {
	baseModel
	Email           string        `gorm:"not null;unique" binding:"required" json:"email"`
	Username        string        `gorm:"not null;unique" binding:"required" json:"username"`
	Age             int           `gorm:"not null"  binding:"required" json:"age"`
	Password        string        `gorm:"not null"  binding:"required" json:"password"`
	ProfileImageURL string        `json:"profile_imgae_url"`
	Photos          []Photo       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Comments        []Comment     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	SocialMedias    []SocialMedia `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
