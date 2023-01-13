package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthorClaims struct {
	jwt.StandardClaims
	Name string
}

type JWTManager struct {
	secretKey string
	duration  time.Duration
}

func (m *JWTManager) Generate(authorName string) (string, error) {
	claims := &AuthorClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(m.duration).Unix(),
		},
		Name: authorName,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(m.secretKey))
}

func (i *JWTManager) Verify(token string) (*AuthorClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(
		token,
		&AuthorClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(i.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := parsedToken.Claims.(*AuthorClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
