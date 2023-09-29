package repositories

import (
	appErr "github.com/BetterToPractice/go-echo-setup/errors"
	"github.com/BetterToPractice/go-echo-setup/lib"
	"github.com/BetterToPractice/go-echo-setup/models"
	"github.com/pkg/errors"
)

type PostRepository struct {
	db lib.Database
}

func NewPostRepository(db lib.Database) PostRepository {
	return PostRepository{
		db: db,
	}
}

func (r PostRepository) Query(params *models.PostQueryParams) (*models.PostPaginationResult, error) {
	db := r.db.ORM.Preload("User").Model(&models.Posts{})
	list := make(models.Posts, 0)

	pagination, err := QueryPagination(db, params.PaginationParam, &list)
	if err != nil {
		return nil, err
	}

	qr := &models.PostPaginationResult{
		Pagination: pagination,
		List:       list,
	}

	return qr, nil
}

func (r PostRepository) Get(id string) (*models.Post, error) {
	post := new(models.Post)

	if ok, err := QueryOne(r.db.ORM.Preload("User").Model(post).Where("id = ?", id), post); err != nil {
		return nil, err
	} else if !ok {
		return nil, appErr.RecordNotFound
	}

	return post, nil
}

func (r PostRepository) Create(post *models.Post) error {
	if err := r.db.ORM.Model(post).Create(post).Error; err != nil {
		return err
	}
	return nil
}

func (r PostRepository) Update(post *models.Post) error {
	if err := r.db.ORM.Model(post).Updates(post).Error; err != nil {
		return err
	}
	return nil
}

func (r PostRepository) Delete(post *models.Post) error {
	if err := r.db.ORM.Model(post).Where("id = ?", post.ID).Delete(post).Error; err != nil {
		return errors.New("invalid, problem with internal")
	}
	return nil
}
