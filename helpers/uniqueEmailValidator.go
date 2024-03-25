package helpers

import (
	"go-learn/models"
	"strings"

	"gorm.io/gorm"
)

func IsUniqueEmail(db *gorm.DB, email string) bool {
	var count int64
	db.Model(&models.User{}).Where("LOWER(email) = ?", strings.ToLower(email)).Count(&count)
	return count == 0
}
