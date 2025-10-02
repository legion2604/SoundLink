package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func CreateAccessToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(3 * time.Hour).Unix() // 3 часа

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	acceessSecret := []byte(os.Getenv("SECRET_ACCESS_KEY"))
	return token.SignedString(acceessSecret)
}
func CreateRefreshToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(30 * 24 * time.Hour).Unix() // 30 дней
	refreshSecret := []byte(os.Getenv("SECRET_REFRESH_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshSecret)
}
