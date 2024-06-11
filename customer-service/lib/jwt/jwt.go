package jwt

import (
	"customer-service/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaimbs struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(UserId string) (string, error) {
	cfg := config.NewConfig()
	var jwtKey = []byte(cfg.JWT_SECRET)
	claims := JwtClaimbs{
		UserId: UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*JwtClaimbs, error) {

	cfg := config.NewConfig()
	var jwtKey = []byte(cfg.JWT_SECRET)
	claims := &JwtClaimbs{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
