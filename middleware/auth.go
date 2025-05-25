package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// generate JWT token
func GenerateJWT(userID uint, role string) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", nil
	}

	return tokenString, nil

}

