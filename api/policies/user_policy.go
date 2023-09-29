package policies

import (
	"errors"
	"github.com/BetterToPractice/go-echo-setup/models"
)

type UserPolicy struct {
}

func NewUserPolicy() UserPolicy {
	return UserPolicy{}
}

func (p UserPolicy) CanDelete(loggedInUser *models.User, user *models.User) (bool, error) {
	if loggedInUser == nil || loggedInUser.ID != user.ID {
		return false, errors.New("unauthorized")
	}
	return true, nil
}
