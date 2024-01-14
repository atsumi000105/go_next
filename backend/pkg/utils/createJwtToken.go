package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type GenerateJWTData struct{
	ID 		uint
	Email 	string
   Date  	time.Time
}

func GenerateJWT(data GenerateJWTData) (string, error) {

    mySigningKey := []byte(GetEnvVariable("JWT_SECRET_KEY"))

    claims := jwt.MapClaims{
		"id":     data.ID,
		"email":  data.Email,
		"date":   time.Now().UTC(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

    return tokenString, nil
}