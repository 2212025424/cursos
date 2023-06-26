package invoice

import (
	"github.com/2212025424/go-db/pkg/invoiceheader"
	"github.com/2212025424/go-db/pkg/invoiceitem"
)

// Model of invoice
type Model struct {
	Header *invoiceheader.Model
	Items  invoiceitem.Models
}

// Storage is an interface that must implement a db storage
type Storage interface {
	Create(*Model) error
}

// Service of invoice
type Service struct {
	storage Storage
}

// Return a pointer of service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Create a new Invoice
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
