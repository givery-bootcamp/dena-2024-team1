package repository

import (
	"myapp/internal/entity"
	"net/http"
)

type UserRepository interface {
	GetAll() ([]entity.User, error)
	CreateUser(username, password string) error
	GetUserByUsername(username string) (entity.User, error)
	SaveSession(r *http.Request, w http.ResponseWriter, user entity.User) error
	GetSessionUser(r *http.Request) (entity.User, error)
	DeleteSessionUser(r *http.Request, w http.ResponseWriter) error
	GetUserPassword(username string) (string, error)
}
