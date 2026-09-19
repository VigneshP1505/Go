package service

import (
	"example.com/first-app/rest/models"
	"example.com/first-app/rest/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	return s.repo.CreateUser(user)
}

//If a struct implements all the methods, it satisfies the interface
