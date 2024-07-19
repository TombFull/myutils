package myutils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	UserID   int    `json:"userID"`
	jwt.RegisteredClaims
}

var mySigningKey = []byte("my_secret_key")

// GenerateJWT 生成Token
func GenerateJWT(username string, id int) (string, error) {
	claims := MyClaims{
		Username: username,
		UserID:   id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Issuer:    "myApp",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT 验证Token
func ValidateJWT(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
