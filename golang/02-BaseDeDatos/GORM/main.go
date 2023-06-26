package main

import (
	"github.com/2212025424/gorm/model"
	"github.com/2212025424/gorm/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	invoice := model.InvoiceHeader{
		Client: "Jose Enrque",
		InvoiceItems: []model.InvoiceItem{
			{ProductID: 2},
			{ProductID: 3},
		},
	}

	storage.DB().Create(&invoice)

	/*myProduct := model.Product{}

	myProduct.ID = 1

	storage.DB().Unscoped().Delete(&myProduct)*/

	/*products := make([]model.Product, 0)

	storage.DB().Find(&products)

	for _, product := range products {
		fmt.Printf("%d - %s\n", product.ID, product.Name)
	}*/

}
