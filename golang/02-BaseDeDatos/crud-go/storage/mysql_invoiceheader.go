package storage

import (
	"database/sql"
	"fmt"

	"github.com/2212025424/go-db/pkg/invoiceheader"
)

const (
	mySQLMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_header (
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		client VARCHAR(25) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP
	)`

	mySQLCreateInvoiceHeader = `INSERT INTO invoice_header (client) VALUES (?)`
)

// MySQLInvoiceHeader used for work with MYSQL
type MySQLInvoiceHeader struct {
	db *sql.DB
}

// NewMySQLInvoiceHeader return a new pointer of MySQLInvoiceHeader
func NewMySQLInvoiceHeader(db *sql.DB) *MySQLInvoiceHeader {
	return &MySQLInvoiceHeader{db}
}

// Migrate implement the interface InvoiceHeader.Storage
func (p *MySQLInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(mySQLMigrateInvoiceHeader)

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
func (p *MySQLInvoiceHeader) CreateTx(tx *sql.Tx, m *invoiceheader.Model) error {
	stmt, err := tx.Prepare(mySQLCreateInvoiceHeader)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(m.Client)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	m.ID = uint(id)

	return nil
}
