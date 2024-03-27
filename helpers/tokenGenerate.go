package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenResult, err := token.SignedString([]byte(os.Getenv("NOT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenResult, nil
}
