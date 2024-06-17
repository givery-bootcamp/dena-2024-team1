package repository

import (
	"myapp/internal/entity"
)

type UserRepository interface {
	GetAll() ([]entity.User, error)
	CreateUser(username, password string) error
}
