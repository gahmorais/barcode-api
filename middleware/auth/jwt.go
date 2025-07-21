package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokenJwt() string {
	tokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": "admin",
		"nbf":  "",
	}).
		fmt.Println("Criando token jwt: ", tokenJwt)
}
