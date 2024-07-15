package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, id string) (string, error) {
	expiration := time.Second * time.Duration(3600*24)
	token :=  jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"expiredAt": time.Now().Add(expiration).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil{
		return "", err
	}
	return tokenString, nil
}