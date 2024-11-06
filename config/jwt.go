package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("jwt-key")

type JWTClaim struct {
	Email string
	jwt.RegisteredClaims
}
