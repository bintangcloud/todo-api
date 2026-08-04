package utils

import (
	"os"
	"time"
	"todo-api/models"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey []byte

func LoadJWTSecret() {
	SecretKey = []byte(os.Getenv("JWT_SECRET"))
}

func GenerateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": user.ID,
			"email":   user.Email,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		},
	)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
