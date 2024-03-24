package helpers

import (
	"go-learn/models"
	"net/url"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func isUniqueEmail(db *gorm.DB, email string) bool {
	var count int64
	db.Model(&models.User{}).Where("LOWER(email) = ?", strings.ToLower(email)).Count(&count)
	return count == 0
}
func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[\w\.-]+@([\w-]+\.)+[\w-]{2,4}$`)
	return emailRegex.MatchString(email)
}
func isUniqueUsername(db *gorm.DB, username string) bool {
	var count int64
	db.Model(&models.User{}).Where("LOWER(username) = ?", strings.ToLower(username)).Count(&count)
	return count == 0
}
func isValidURL(uri string) bool {
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		return false
	}
	return true
}
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
