package storage

import (
	"database/sql"
	"fmt"

	"github.com/2212025424/go-db/pkg/product"
)

const (
	mySQLMigrateProduct = `CREATE TABLE IF NOT EXISTS product (
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price DECIMAL(7, 2),
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP
	)`

	mySQLCreateProduct = `INSERT INTO product (
		name,observations,price,created_at
	) VALUES (
		?, ?, ?, ?
	)`

	mySQLGetAllProduct = `SELECT id, name, observations, price, created_at, updated_at FROM product`

	mySQLGetProductById = mySQLGetAllProduct + ` WHERE id = ?`

	mySQLUpdateProduct = `UPDATE product SET name = ?, observations = ?, price = ?, updated_at = ? WHERE id = ?`

	mySQLDeleteProduct = `DELETE FROM product WHERE id = ?`
)

// MySQLProduct used for work with MySQL
type MySQLProduct struct {
	db *sql.DB
}

// newMySQLProduct return a new pointer of MySQLProduct
func newMySQLProduct(db *sql.DB) *MySQLProduct {
	return &MySQLProduct{db}
}

// Migrate implement the interface product.Storage
func (p *MySQLProduct) Migrate() error {
	stmt, err := p.db.Prepare(mySQLMigrateProduct)

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
func (p *MySQLProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(mySQLCreateProduct)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(m.Name, stringToNull(m.Observations), m.Price, m.CreatedAt)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	m.ID = uint(id)

	fmt.Printf("Se ha agregado el producto con id %d", m.ID)

	return nil
}

// GetAll implement the interface product.Storage
func (p *MySQLProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(mySQLGetAllProduct)

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
func (p *MySQLProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(mySQLGetProductById)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))
}

// Update implement th interface product.Storage
func (p *MySQLProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(mySQLUpdateProduct)

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
func (p *MySQLProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(mySQLDeleteProduct)

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
