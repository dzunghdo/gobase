package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"

	"gobase/models/entities"
	"gobase/services/auth/dto"
)

func CreateJWTToken(user *entities.User, key []byte) (string, dto.Claims, error) {
	claims := &dto.Claims{
		Username: user.Username,
		Email:    user.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", dto.Claims{}, err
	}
	return tokenString, *claims, nil
}

func ExtractClaimsFromToken(token string, key []byte) (*dto.Claims, error) {
	var (
		claims = &dto.Claims{}
		err    error
	)
	tokenClaims, err := jwt.ParseWithClaims(token, claims,
		func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})
	if err != nil {
		return nil, errors.New("Failed to parse JWT token")
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*dto.Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return claims, nil
}
