package repositories

import (
	"errors"
	"github.com/BetterToPractice/go-echo-setup/api/dto"
	appErr "github.com/BetterToPractice/go-echo-setup/errors"
	"github.com/BetterToPractice/go-echo-setup/lib"
	"github.com/BetterToPractice/go-echo-setup/models"
)

type UserRepository struct {
	db lib.Database
}

func NewUserRepository(db lib.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r UserRepository) Query(params *dto.UserQueryParam) (*dto.UserPaginationResponse, error) {
	db := r.db.ORM.Preload("Profile").Model(models.Users{})

	var list models.Users
	pagination, err := QueryPagination(db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.Join(appErr.DatabaseInternalError, err)
	}

	qr := &dto.UserPaginationResponse{Pagination: pagination}
	qr.Serializer(&list)

	return qr, nil
}

func (r UserRepository) GetByUsername(username string) (*models.User, error) {
	user := new(models.User)

	if ok, err := QueryOne(r.db.ORM.Preload("Profile").Model(user).Where("username = ?", username), user); err != nil {
		return nil, errors.Join(appErr.DatabaseInternalError, err)
	} else if !ok {
		return nil, appErr.DatabaseRecordNotFound
	}

	return user, nil
}

func (r UserRepository) Create(user *models.User) error {
	if err := r.db.ORM.Model(user).Create(user).Error; err != nil {
		return errors.Join(appErr.DatabaseInternalError, err)
	}
	return nil
}

func (r UserRepository) Delete(user *models.User) error {
	if err := r.db.ORM.Model(user).Where("username = ?", user.ID).Delete(user).Error; err != nil {
		return errors.Join(appErr.DatabaseInternalError, err)
	}

	return nil
}
