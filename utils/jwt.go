package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("supersecretkey")

func GenerateToken(email string) (string, error){
	
	//var jwtSecret = []byte("supersecretkey")

	claims := jwt.MapClaims{
		"email" : email,
		"exp" : time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret);

	return tokenString, err
}
