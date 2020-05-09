package service

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"github.com/rosspatil/go-kit-example/pb"
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
func (s *Service) RegisterEmployee(ctx context.Context, employee pb.Employee) (string, error) {
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
func (s *Service) GetEmployeeDetails(ctx context.Context, ID string) (*pb.Employee, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetEmployeeDetails")
	defer span.Finish()
	return storage.GetClient().Get(ctx, ID)
}

// DeleteEmployee -...
func (s *Service) DeleteEmployee(ctx context.Context, ID string) error {
	return storage.GetClient().Delete(ctx, ID)
}
