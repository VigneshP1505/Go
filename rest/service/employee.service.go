package service

import (
	"example.com/first-app/rest/models"
	"example.com/first-app/rest/repository"
)

type EmployeeService struct {
	Repo repository.EmployeeRepository
}

func (s *EmployeeService) GetAll() ([]models.Employee, error) {
	return s.Repo.GetAll()
}

func (s *EmployeeService) CreateEmployee(e models.Employee) (models.Employee, error) {
	return s.Repo.CreateEmployee(e)
}
