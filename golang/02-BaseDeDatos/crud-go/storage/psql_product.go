package storage

import (
	"database/sql"
	"fmt"

	"github.com/2212025424/go-db/pkg/product"
)

type scanner interface {
	Scan(dest ...interface{}) error
}

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS product (
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price REAL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT product_id_pk PRIMARY KEY (id)
	)`

	psqlCreateProduct = `INSERT INTO product (
		name,observations,price,created_at
	) VALUES (
		$1, $2, $3, $4
	) RETURNING id`

	psqlGetAllProduct = `SELECT id, name, observations, price, created_at, updated_at FROM product`

	psqlGetProductById = psqlGetAllProduct + ` WHERE id = $1`

	psqlUpdateProduct = `UPDATE product SET name = $1, observations = $2, price = $3, updated_at = $4 WHERE id = $5`

	psqlDeleteProduct = `DELETE FROM product WHERE id = $1`
)

// PsqlProduct used for work with postgres
type PsqlProduct struct {
	db *sql.DB
}

// newPsqlProduct return a new pointer of PsqlProduct
func newPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

// Migrate implement the interface product.Storage
func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		return err
	}

	fmt.Println("Migracion de producto ejecutada correctamente...")
	return nil
}

// Create implement the interfaces product.Storage
func (p *PsqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)

	if err != nil {
		return err
	}

	defer stmt.Close()

	err = stmt.QueryRow(m.Name, stringToNull(m.Observations), m.Price, m.CreatedAt).Scan(&m.ID)

	if err != nil {
		return err
	}

	fmt.Println("Se ha agregado el producto")

	return nil
}

// GetAll implement the interface product.Storage
func (p *PsqlProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(psqlGetAllProduct)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	ms := make(product.Models, 0)
	for rows.Next() {
		m, err := scanRowProduct(rows)

		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ms, nil
}

// GetById implement the interface product.storage
func (p *PsqlProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(psqlGetProductById)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))
}

// Update implement th interface product.Storage
func (p *PsqlProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlUpdateProduct)

	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		timeToNull(m.UpdatedAt),
		m.ID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no existe el producto con ID: %d", m.ID)
	}

	fmt.Println("Se ha actualizado el producto")

	return nil
}

// Delete implements the interface product.Storage
func (p *PsqlProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(psqlDeleteProduct)

	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(id)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no existe el producto con ID: %d", id)
	}

	fmt.Println("Se ha eliminado el producto")

	return nil
}

func scanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}

	observationsNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&observationsNull,
		&m.Price,
		&m.CreatedAt,
		&updatedAtNull,
	)

	if err != nil {
		return nil, err
	}

	m.Observations = observationsNull.String
	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}
