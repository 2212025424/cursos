package main

import (
	"fmt"
	"log"

	"github.com/2212025424/go-db/pkg/product"
	"github.com/2212025424/go-db/storage"
)

func main() {

	driver := storage.MySQL

	storage.New(driver)

	myStorage, err := storage.DAOProduct(driver)

	if err != nil {
		log.Fatalf("DAOProduct: %v", err)
	}

	serviceProduct := product.NewService(myStorage)

	ms, err := serviceProduct.GetAll()

	if err != nil {
		log.Fatalf("Product.GetAll: %v", err)
	}

	fmt.Println(ms)
}

/*storageProduct := storage.NewPsqlProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

if err := serviceProduct.Migrate(); err != nil {
	log.Fatalf("product.Migrate: %v", err)
}

storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)

if err := serviceInvoiceHeader.Migrate(); err != nil {
	log.Fatalf("InvoiceHeader.Migrate: %v", err)
}

storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)

if err := serviceInvoiceItem.Migrate(); err != nil {
	log.Fatalf("InvoiceItem.Migrate: %v", err)
}

storageProduct := storage.NewPsqlProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

m := &product.Model{
	Name:         "Curso DB con GO",
	Observations: "On fire",
	Price:        70,
}

if err := serviceProduct.Create(m); err != nil {
	log.Fatalf("product.Create: %v", err)
}


storageProduct := storage.NewPsqlProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

ms, err := serviceProduct.GetAll()

if err != nil {
	log.Fatalf("product.GetAll: %v", err)
}


storageHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
storageItems := storage.NewPsqlInvoiceItem(storage.Pool())
storageInvoice := storage.NewPsqlInvoice(storage.Pool(), storageHeader, storageItems)

m := &invoice.Model{
	Header: &invoiceheader.Model{
		Client: "Fernanda",
	},
	Items: invoiceitem.Models{
		&invoiceitem.Model{ProductID: 1},
		&invoiceitem.Model{ProductID: 2},
	},
}

serviceInvoice := invoice.NewService(storageInvoice)
if err := serviceInvoice.Create(m); err != nil {
	log.Fatalf("Invoice.create: %v", err)
}


*/
