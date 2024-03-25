package helpers

import (
	"go-learn/models"
	"strings"

	"gorm.io/gorm"
)

func IsUniqueUsername(db *gorm.DB, username string) bool {
	var count int64
	db.Model(&models.User{}).Where("LOWER(username) = ?", strings.ToLower(username)).Count(&count)
	return count == 0
}
