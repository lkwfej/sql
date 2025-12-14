package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, username string) (string, error) {
	return GenerateTokenWithSecret(userID, username, []byte("change_this_secret"), 24*time.Hour)
}

// GenerateTokenWithSecret creates a JWT token.
// secret and ttl are injected from main/config (avoid hardcoding secrets in code).
func GenerateTokenWithSecret(userID uint, username string, secret []byte, ttl time.Duration) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ParseToken(tokenStr string) (*Claims, error) {
	return ParseTokenWithSecret(tokenStr, []byte("change_this_secret"))
}

// ParseTokenWithSecret validates a JWT token and returns claims.
func ParseTokenWithSecret(tokenStr string, secret []byte) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		// Ensure expected signing method (avoid "alg=none" style issues)
		if token.Method == nil || token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token is invalid")
	}
	return claims, nil
}
