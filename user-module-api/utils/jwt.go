package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

/*
GenerateToken
User login ke baad JWT token banata hai
*/
func GenerateToken(userID string) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("JWT_SECRET")

	return token.SignedString([]byte(secret))
}

/*
ValidateToken
Middleware me token verify karta hai
*/
func ValidateToken(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		secret = "your_jwt_secret_here"
	}

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
}
