package invoiceheader

import (
	"database/sql"
	"time"
)

// Model of invoiceheader
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Storage interface that must implement a db storage
type Storage interface {
	Migrate() error
	CreateTx(*sql.Tx, *Model) error
}

// Service of InvoiceHeader
type Service struct {
	storage Storage
}

// NewService return a pointer of service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used for migrate InvoiceHeader
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
