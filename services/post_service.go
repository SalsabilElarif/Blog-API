package services

import (
	"blog-api/models"
	"blog-api/repository"
)

type PostService struct {
	Repo repository.PostRepo
}

func (s *PostService) CreatePost(post *models.Post) error {
	err := s.Repo .CreatePost(post)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostService) GetPostByID(id uint) (models.Post, error) {
	post, err := s.Repo.GetPostByID(id)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}
