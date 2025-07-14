package repository

import (
	"blog-api/database"
	"blog-api/models"
)

type PostRepository interface {
	CreatePost(post *models.Post) error
}

type PostRepo struct{}

func (r *PostRepo) CreatePost(post *models.Post) error {
	err := database.DB.Create(post).Error
	if err != nil {
		return err
	} 
	return nil 
}

func (r *PostRepo) GetPostByID(id uint) (models.Post, error) {
	var post models.Post
	err := database.DB.Preload("User").Where("id = ?", id).Find(&post).Error
	if err != nil {
		return models.Post{}, err 
	}

	return post, nil
}