package dto

import "github.com/BetterToPractice/go-echo-setup/models"

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterResponse struct {
	Username string
	Email    string
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Access string `json:"access"`
}

func (r *RegisterResponse) Serializer(user *models.User) {
	r.Username = user.Username
	r.Email = user.Email
}

func (r *LoginResponse) Serializer(access string) {
	r.Access = access
}
