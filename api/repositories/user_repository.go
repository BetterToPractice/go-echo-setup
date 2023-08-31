package repositories

import (
	"errors"
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
	db := r.db.ORM.Model(&models.Users{})
	list := make(models.Users, 0)

	pagination, err := QueryPagination(db, params.PaginationParam, &list)
	if err != nil {
		return nil, err
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
		return nil, errors.New("error database")
	} else if !ok {
		return nil, errors.New("not found")
	}

	return user, nil
}
