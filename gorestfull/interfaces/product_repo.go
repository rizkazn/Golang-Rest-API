package interfaces

import "github.com/rizkazn/gorestfull/models"

type ProdRepository interface {
	FindAll() (*models.Products, error)
	Save(product *models.Product) error
	Update(product *models.Product, id string) error
	Remove(id string) error
	SearchProductBytName(name string) (*models.Products, error)
	OrderProdByName() (*models.Products, error)
	OrderProdByCategory() (*models.Products, error)
	OrderProdByNewest() (*models.Products, error)
	OrderProdByPrice() (*models.Products, error)
}
