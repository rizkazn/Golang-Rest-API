package interfaces

import "github.com/rizkazn/gorestfull/models"

type UserRepository interface {
	FindAll() (*models.Users, error)
	Save(history *models.User) error
	GetPassword(username string) (string, error)
	GetUsername(username string) (string, error)
	GetEmail(username string) (string, error)
	GetRole(username string) (string, error)
}
