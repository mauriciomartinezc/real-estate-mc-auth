package domain

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type JWTCustomClaims struct {
	UserID uuid.UUID `json:"id"`
	Email  string    `json:"email"`
	jwt.StandardClaims
}
