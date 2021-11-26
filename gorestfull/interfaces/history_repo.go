package interfaces

import "github.com/rizkazn/gorestfull/models"

type HistRepository interface {
	FindAll() (*models.Histories, error)
	Save(history *models.History) error
	Update(history *models.History, id string) error
	Remove(id string) error
}
