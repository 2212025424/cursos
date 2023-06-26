package storage

import (
	"database/sql"
	"fmt"

	"github.com/2212025424/go-db/pkg/invoiceheader"
)

const (
	psqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_header (
		id SERIAL NOT NULL,
		client VARCHAR(25) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_header_id_pk PRIMARY KEY (id)
	)`

	psqlCreateInvoiceHeader = `INSERT INTO invoice_header (client) VALUES ($1) RETURNING id, created_at`
)

// PsqlInvoiceHeader used for work with postgres
type PsqlInvoiceHeader struct {
	db *sql.DB
}

// NewPsqlInvoiceHeader return a new pointer of PsqlInvoiceHeader
func NewPsqlInvoiceHeader(db *sql.DB) *PsqlInvoiceHeader {
	return &PsqlInvoiceHeader{db}
}

// Migrate implement the interface InvoiceHeader.Storage
func (p *PsqlInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceHeader)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		return err
	}

	fmt.Println("Migracion de InvoiceHeader ejecutada correctamente...")
	return nil
}

// CreateTx implement the interface invoiceHeader.Storage
func (p *PsqlInvoiceHeader) CreateTx(tx *sql.Tx, m *invoiceheader.Model) error {
	stmt, err := tx.Prepare(psqlCreateInvoiceHeader)

	if err != nil {
		return err
	}

	defer stmt.Close()

	return stmt.QueryRow(m.Client).Scan(&m.ID, &m.CreatedAt)
}
