package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userID int64, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = userID
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}
