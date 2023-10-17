package dto

import (
	"github.com/BetterToPractice/go-echo-setup/models"
)

type UserQueryParam struct {
	PaginationParam
}

type ProfileResponse struct {
	PhoneNumber string `json:"phone_number,omitempty"`
	Gender      string `json:"gender,omitempty"`
}

type UserResponse struct {
	Username string          `json:"username"`
	Email    string          `json:"email"`
	Profile  ProfileResponse `json:"profile"`
}

func (r *UserResponse) Serializer(user *models.User) {
	r.Username = user.Username
	r.Email = user.Email
	r.Profile = ProfileResponse{
		PhoneNumber: user.Profile.PhoneNumber,
		Gender:      user.Profile.Gender,
	}
}

type UserPaginationResponse struct {
	List       []UserResponse `json:"list"`
	Pagination *Pagination    `json:"pagination"`
}

func (p *UserPaginationResponse) Serializer(users *models.Users) {
	var list []UserResponse
	for _, user := range *users {
		u := UserResponse{}
		u.Serializer(&user)
		list = append(list, u)
	}
	p.List = list
}
