package repositories

import (
	"github.com/BetterToPractice/go-echo-setup/api/dto"
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

func (r PostRepository) Query(params *dto.PostQueryParam) (*dto.PostPaginationResponse, error) {
	posts := new(models.Posts)
	db := r.db.ORM.Preload("User").Model(posts)

	var list models.Posts
	pagination, err := QueryPagination(db, params.PaginationParam, &list)
	if err != nil {
		return nil, err
	}

	qr := &dto.PostPaginationResponse{Pagination: pagination}
	qr.Serializer(&list)

	return qr, nil
}

func (r PostRepository) Get(id string) (*models.Post, *dto.PostResponse, error) {
	post := new(models.Post)

	if ok, err := QueryOne(r.db.ORM.Preload("User").Model(post).Where("id = ?", id), post); err != nil {
		return nil, nil, err
	} else if !ok {
		return nil, nil, appErr.RecordNotFound
	}

	resp := &dto.PostResponse{}
	resp.Serializer(post)

	return post, resp, nil
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
