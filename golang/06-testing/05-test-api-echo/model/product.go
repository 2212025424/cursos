package model

import (
	"errors"
)

var (
	ErrIDNotFound = errors.New("no hay elemento asociado a ese identificador")
)

// Product of product
type Product struct {
	ID    uint
	Name  string
	Price int
}

// Products slice of Product
type Products []*Product

// Storage interface that must implement a db storage
type Storage interface {
	Migrate() error
	Create(*Product) error
	GetAll() (Products, error)
	GetByID(uint) (*Product, error)
	Update(*Product) error
	Delete(uint) error
}

//Service of Product
type Service struct {
	storage Storage
}

// NewService return a pointer of service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used for migrate product
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}

// Create is used for create a product
func (s *Service) Create(m *Product) error {
	return s.storage.Create(m)
}

// GetAll is used for get all products
func (s *Service) GetAll() (Products, error) {
	return s.storage.GetAll()
}

// GetByID is used for get a product
func (s *Service) GetByID(id uint) (*Product, error) {
	return s.storage.GetByID(id)
}

// Update is used for update a product
func (s *Service) Update(m *Product) error {
	if m.ID == 0 {
		return ErrIDNotFound
	}

	return s.storage.Update(m)
}

// Delete is used for delete a product
func (s *Service) Delete(id uint) error {
	return s.storage.Delete(id)
}
