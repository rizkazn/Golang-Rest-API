package interfaces

import "github.com/rizkazn/gorestfull/models"

type CatRepository interface {
	FindAll() (*models.Categories, error)
	Save(category *models.Category) error
	Update(category *models.Category, id string) error
	Remove(id string) error
}
