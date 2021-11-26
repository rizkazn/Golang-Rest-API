package models

import (
	"time"
)

type Category struct {
	Id             int       `json:"id,omitempty"`
	Category_name  string    `json:"category_name"`
	Category_image string    `json:"category_image"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Categories []Category

func CreateCategory() *Category {
	return &Category{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
