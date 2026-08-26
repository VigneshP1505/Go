package repository

import "example.com/first-app/rest/models"

type EmployeeRepository interface {
	GetAll() ([]models.Employee, error)
	CreateEmployee(e models.Employee) (models.Employee, error)
}
