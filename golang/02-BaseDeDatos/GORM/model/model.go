package model

//https://app.ed.team/cursos/database-go/04/04
import (
	"gorm.io/gorm"
)

// Model of product
type Product struct {
	gorm.Model
	Name         string        `gorm:"varchar(100); not null"`
	Observations *string       `gorm:"varchar(100)"`
	Price        int           `gorm:"not null"`
	InvoiceItems []InvoiceItem `gorm:"not null"`
}

// Model of invoiceheader
type InvoiceHeader struct {
	gorm.Model
	Client       string        `gorm:"varchar(100); not null"`
	InvoiceItems []InvoiceItem `gorm:"not null"`
}

// Model of Invoiceitem
type InvoiceItem struct {
	gorm.Model
	InvoiceHeaderID uint
	ProductID       uint
}
