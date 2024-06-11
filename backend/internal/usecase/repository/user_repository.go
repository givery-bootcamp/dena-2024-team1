package repository

import (
	"myapp/internal/entity"
)

type UserRepository interface {
	GetAll() ([]entity.User, error)
}
