package services

import (
	"github.com/BetterToPractice/go-echo-setup/api/repositories"
	"github.com/BetterToPractice/go-echo-setup/models"
	"github.com/BetterToPractice/go-echo-setup/models/dto"
)

type PostService struct {
	postRepository repositories.PostRepository
}

func NewPostService(postRepository repositories.PostRepository) PostService {
	return PostService{
		postRepository: postRepository,
	}
}

func (s PostService) Query(params *models.PostQueryParams) (*models.PostPaginationResult, error) {
	return s.postRepository.Query(params)
}

func (s PostService) Get(id string) (*models.Post, error) {
	return s.postRepository.Get(id)
}

func (s PostService) Create(params *dto.PostRequest, user *models.User) (*dto.PostResponse, error) {
	post := &models.Post{Title: params.Title, Body: params.Body, UserID: user.ID}
	if err := s.postRepository.Create(post); err != nil {
		return nil, err
	}

	return &dto.PostResponse{Title: post.Title, Body: post.Body}, nil
}

func (s PostService) Delete(id string) error {
	_, err := s.postRepository.Get(id)
	if err != nil {
		return err
	}

	return s.postRepository.Delete(id)
}
