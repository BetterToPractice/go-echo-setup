package dto

import "github.com/golang-jwt/jwt/v5"

type JWTResponse struct {
	Access string `json:"access" validate:"required"`
}

type JWTClaims struct {
	ID       uint
	Username string
	jwt.RegisteredClaims
}
