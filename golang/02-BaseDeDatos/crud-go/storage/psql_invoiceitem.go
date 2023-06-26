package storage

import (
	"database/sql"
	"fmt"

	"github.com/2212025424/go-db/pkg/invoiceitem"
)

const (
	psqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_item (
		id SERIAL NOT NULL,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_item_id_pk PRIMARY KEY (id),
		CONSTRAINT invoice_item_invoice_header_id_fk FOREIGN KEY 
		(invoice_header_id) REFERENCES invoice_header (id)
			ON UPDATE RESTRICT
			ON DELETE RESTRICT,
		CONSTRAINT invoice_item_product_id_fk FOREIGN KEY 
		(product_id) REFERENCES product (id)
			ON UPDATE RESTRICT
			ON DELETE RESTRICT
	)`

	psqlCreateInvoiceItem = `INSERT INTO invoice_item (invoice_header_id, product_id) VALUES ($1, $2) RETURNING id, created_at`
)

// PsqlInvoiceItem used for work with postgres
type PsqlInvoiceItem struct {
	db *sql.DB
}

// NewPsqlInvoiceItem return a new pointer of PsqlInvoiceItem
func NewPsqlInvoiceItem(db *sql.DB) *PsqlInvoiceItem {
	return &PsqlInvoiceItem{db}
}

// Migrate implement the interface InvoiceItem.Storage
func (p *PsqlInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceItem)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		return err
	}

	fmt.Println("Migracion de InvoiceItem ejecutada correctamente...")
	return nil
}

// CreateTx implement the interface invoiceItem.Storage
func (p *PsqlInvoiceItem) CreateTx(tx *sql.Tx, headerID uint, ms invoiceitem.Models) error {
	stmt, err := tx.Prepare(psqlCreateInvoiceItem)

	if err != nil {
		return err
	}

	defer stmt.Close()

	for _, item := range ms {
		err = stmt.QueryRow(headerID, item.ProductID).Scan(
			&item.ID,
			&item.CreatedAt,
		)

		if err != nil {
			return err
		}
	}

	return nil
}
