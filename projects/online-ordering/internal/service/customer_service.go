package service

import "github.com/vignesh/online-ordering/internal/repository"

type CustomerService struct {
	repo repository.CustomerRepository
}
