package storage

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/rosspatil/go-kit-example/models"
)

// Iface - storage interface
type Iface interface {
	Get(ctx context.Context, ID string) (*models.Employee, error)
	Create(ctx context.Context, employee models.Employee) (string, error)
	Update(ctx context.Context, ID string, employee models.Employee) error
	Delete(ctx context.Context, ID string) error
}

// Storage - this is sti implementor
type Storage struct {
	db *DB
	Iface
}

var client *Storage
var once sync.Once

func GetClient() *Storage {
	once.Do(func() {
		client = &Storage{db: NewClient()}
	})
	return client
}

func (s *Storage) Create(ctx context.Context, employee models.Employee) (string, error) {
	employee.ID = uuid.New().String()
	err := s.db.Set(ctx, employee.ID, employee)
	if err != nil {
		return "", err
	}
	return employee.ID, nil
}

func (s *Storage) Update(ctx context.Context, ID string, employee models.Employee) error {
	err := s.db.Set(ctx, ID, employee)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Delete(ctx context.Context, ID string) error {
	err := s.db.Delete(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Get(ctx context.Context, ID string) (*models.Employee, error) {
	data, err := s.db.Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	employee := data.(models.Employee)
	return &employee, nil
}
