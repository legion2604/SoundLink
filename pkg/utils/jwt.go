package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var accessSecret = []byte("secret_access")

func CreateAccessToken(userPass string) (string, error) {
	claims := jwt.MapClaims{}
	claims["password"] = userPass
	claims["exp"] = time.Now().Add(122 * time.Minute).Unix() // 2 часа 2 минуты

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(accessSecret)
}

var refreshSecret = []byte("secret_refresh")

func CreateRefreshToken(userPass string) (string, error) {
	claims := jwt.MapClaims{}
	claims["password"] = userPass
	claims["exp"] = time.Now().Add(30 * 24 * time.Hour).Unix() // 30 дней

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshSecret)
}
