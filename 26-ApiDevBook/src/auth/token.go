package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userID int) (string, error) {
	permissions := jwt.MapClaims{
		"userId": userID,
	}

	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte("secret"))
}