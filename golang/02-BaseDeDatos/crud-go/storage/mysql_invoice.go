package storage

import (
	"database/sql"
	"fmt"

	"github.com/2212025424/go-db/pkg/invoice"
	"github.com/2212025424/go-db/pkg/invoiceheader"
	"github.com/2212025424/go-db/pkg/invoiceitem"
)

// MySQLInvoice used for work with postgres - invoice
type MySQLInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	invoiceItems  invoiceitem.Storage
}

// NewMySQLInvoice return a pointer of MySQLInvoice
func NewMySQLInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *MySQLInvoice {
	return &MySQLInvoice{
		db:            db,
		storageHeader: h,
		invoiceItems:  i,
	}
}

// Create implement the interface invoice.Storage
func (p *MySQLInvoice) Create(m *invoice.Model) error {
	tx, err := p.db.Begin()

	if err != nil {
		return err
	}

	if err := p.storageHeader.CreateTx(tx, m.Header); err != nil {
		tx.Rollback()
		return fmt.Errorf("header: %w", err)
	}

	fmt.Printf("Se ha creado la factura con ID: %d \n", m.Header.ID)

	if err := p.invoiceItems.CreateTx(tx, m.Header.ID, m.Items); err != nil {
		tx.Rollback()
		return fmt.Errorf("items: %w", err)
	}

	fmt.Printf("Items creados: %d", len(m.Items))

	return tx.Commit()
}
