package models

import (
	"time"
)

type Product struct {
	Id            int       `json:"id,omitempty"`
	Product_name  string    `json:"product_name"`
	Price         int       `json:"price"`
	Product_image string    `json:"product_image"`
	Category_id   int       `json:"category_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Products []Product

func CreateProduct() *Product {
	return &Product{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
