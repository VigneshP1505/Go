package repository

import "example.com/first-app/rest/models"

type UserRepository interface {
	GetAll() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
}
