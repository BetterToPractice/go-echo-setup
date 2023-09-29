package policies

import (
	"errors"
	"github.com/BetterToPractice/go-echo-setup/models"
)

type PostPolicy struct {
}

func NewPostPolicy() PostPolicy {
	return PostPolicy{}
}

func (p PostPolicy) CanViewList(_ *models.User) (bool, error) {
	return true, nil
}

func (p PostPolicy) CanViewDetail(_ *models.User, _ *models.Post) (bool, error) {
	return true, nil
}

func (p PostPolicy) CanCreate(user *models.User) (bool, error) {
	if user == nil {
		return false, errors.New("unauthorized")
	}
	return true, nil
}

func (p PostPolicy) CanUpdate(user *models.User, post *models.Post) (bool, error) {
	if user == nil || post.UserID != user.ID {
		return false, errors.New("unauthorized")
	}
	return true, nil
}

func (p PostPolicy) CanDelete(user *models.User, post *models.Post) (bool, error) {
	if user == nil || post.UserID != user.ID {
		return false, errors.New("unauthorized")
	}
	return true, nil
}
