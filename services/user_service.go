package services

import (
	"blog-api/models"
	"blog-api/repository"
	"blog-api/utils"
	"errors"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) RegisterUser(user *models.User) error {
	// check if user exists already
	_, err := s.Repo.GetUserByEmail(user.Email)
	if err == nil {
		return errors.New("email already in use")
	}

	// hash the password
	hashpass, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashpass

	// call the create method
	err = s.Repo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}
