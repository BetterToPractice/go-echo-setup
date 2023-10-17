package repositories

import (
	"errors"
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

func (r UserRepository) Query(params *models.UserQueryParams) (*models.UserPaginationResult, error) {
	db := r.db.ORM.Preload("Profile").Model(&models.Users{})
	list := make(models.Users, 0)

	pagination, err := QueryPagination(db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.Join(appErr.DatabaseInternalError, err)
	}

	qr := &models.UserPaginationResult{
		Pagination: pagination,
		List:       list,
	}

	return qr, nil
}

func (r UserRepository) GetByUsername(username string) (*models.User, error) {
	user := new(models.User)

	if ok, err := QueryOne(r.db.ORM.Model(user).Where("username = ?", username), user); err != nil {
		return nil, errors.Join(appErr.DatabaseInternalError, err)
	} else if !ok {
		return nil, appErr.DatabaseRecordNotFound
	}

	return user, nil
}

func (r UserRepository) Delete(user *models.User) error {
	if err := r.db.ORM.Model(user).Where("username = ?", user.ID).Delete(user).Error; err != nil {
		return errors.Join(appErr.DatabaseInternalError, err)
	}

	return nil
}
