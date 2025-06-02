package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GenerateJWT(userID uint, tokenversion string) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"sub":           userID,
		"token_version": tokenversion,
		"exp":           now.Add(time.Hour * 24).Unix(),
		"iat":           now.Unix(),
		"iss":           os.Getenv("APP_NAME"),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}
