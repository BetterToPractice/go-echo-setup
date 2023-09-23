package services

import (
	"github.com/BetterToPractice/go-echo-setup/api/repositories"
	"github.com/BetterToPractice/go-echo-setup/models"
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

func (s PostService) Delete(id string) error {
	_, err := s.postRepository.Get(id)
	if err != nil {
		return err
	}

	return s.postRepository.Delete(id)
}
