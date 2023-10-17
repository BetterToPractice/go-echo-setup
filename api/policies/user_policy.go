package policies

import (
	appError "github.com/BetterToPractice/go-echo-setup/errors"
	"github.com/BetterToPractice/go-echo-setup/models"
)

type UserPolicy struct {
}

func NewUserPolicy() UserPolicy {
	return UserPolicy{}
}

func (p UserPolicy) CanDelete(loggedInUser *models.User, user *models.User) (bool, error) {
	if loggedInUser == nil {
		return false, appError.Unauthorized
	}
	if loggedInUser.ID != user.ID {
		return false, appError.Forbidden
	}
	return true, nil
}
