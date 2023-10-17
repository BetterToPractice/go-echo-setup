package services

import (
	"errors"
	"github.com/BetterToPractice/go-echo-setup/api/dto"
	"github.com/BetterToPractice/go-echo-setup/api/mails"
	"github.com/BetterToPractice/go-echo-setup/api/repositories"
	appError "github.com/BetterToPractice/go-echo-setup/errors"
	"github.com/BetterToPractice/go-echo-setup/models"
)

type PostService struct {
	postRepository repositories.PostRepository
	postMail       mails.PostMail
}

func NewPostService(postRepository repositories.PostRepository, postMail mails.PostMail) PostService {
	return PostService{
		postRepository: postRepository,
		postMail:       postMail,
	}
}

func (s PostService) Query(params *dto.PostQueryParam) (*dto.PostPaginationResponse, error) {
	return s.postRepository.Query(params)
}

func (s PostService) Get(id string) (*models.Post, *dto.PostResponse, error) {
	post, resp, err := s.postRepository.Get(id)
	if err != nil && errors.Is(err, appError.DatabaseRecordNotFound) {
		err = appError.PostNotFound
	}
	return post, resp, err
}

func (s PostService) Create(params *dto.PostRequest, user *models.User) (*dto.PostResponse, error) {
	post := &models.Post{Title: params.Title, Body: params.Body, UserID: user.ID}
	if err := s.postRepository.Create(post); err != nil {
		return nil, err
	}

	s.postMail.CreatePost(user, post)

	resp := &dto.PostResponse{}
	resp.Serializer(post)

	return resp, nil
}

func (s PostService) Update(post *models.Post, params *dto.PostUpdateRequest) (*dto.PostResponse, error) {
	if params.Title != "" {
		post.Title = params.Title
	}
	if params.Body != "" {
		post.Body = params.Body
	}

	if err := s.postRepository.Update(post); err != nil {
		return nil, err
	}

	resp := &dto.PostResponse{}
	resp.Serializer(post)

	return resp, nil
}

func (s PostService) Delete(post *models.Post) error {
	return s.postRepository.Delete(post)
}
