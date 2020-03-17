package service

import (
	"context"

	"github.com/rosspatil/go-kit-example/models"
	"github.com/rosspatil/go-kit-example/storage"
)

// MyService - ...
type MyService interface {
	RegisterEmployee()
	UpdateEmail()
	GetEmployeeDetails()
	DeleteEmployee()
}

// Service - ...
type Service struct {
	MyService
}

// NewService - ...
func NewService() *Service {
	return new(Service)
}

// RegisterEmployee -...
func (s *Service) RegisterEmployee(ctx context.Context, employee models.Employee) (string, error) {
	return storage.GetClient().Create(ctx, employee)
}

// UpdateEmail - ...
func (s *Service) UpdateEmail(ctx context.Context, ID, email string) error {
	e, err := storage.GetClient().Get(ctx, ID)
	if err != nil {
		return err
	}
	e.Email = email
	return storage.GetClient().Update(ctx, ID, *e)
}

// GetEmployeeDetails -...
func (s *Service) GetEmployeeDetails(ctx context.Context, ID string) (*models.Employee, error) {
	return storage.GetClient().Get(ctx, ID)
}

// DeleteEmployee -...
func (s *Service) DeleteEmployee(ctx context.Context, ID string) error {
	return storage.GetClient().Delete(ctx, ID)
}
