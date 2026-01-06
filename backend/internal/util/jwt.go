package util

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID    string `json:"user_id"`
	CompanyID string `json:"company_id"`
	Role      string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(secret, userID, companyID, role string, ttl time.Duration) (string, error) {
	if secret == "" {
		return "", errors.New("JWT secret is empty")
	}

	now := time.Now()

	claims := Claims{
		UserID:    userID,
		CompanyID: companyID,
		Role:      role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ParseToken(secret, tokenString string) (*Claims, error) {
	if secret == "" {
		return nil, errors.New("JWT secret is empty")
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
func GenerateJWT(secret, userID, companyID, role string) (string, error) {
	return GenerateToken(secret, userID, companyID, role, 7*24*time.Hour)
}

func ValidateJWT(secret, tokenString string) (*Claims, error) {
	return ParseToken(secret, tokenString)
}
