package repository

import (
	"context"
	"myapp/internal/entity"

	"github.com/gin-contrib/sessions"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]entity.User, error)
	CreateUser(ctx context.Context, username, password string) (*entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	SaveSession(ctx context.Context, session sessions.Session, user entity.User) error
	GetSessionUser(ctx context.Context, session sessions.Session) (*entity.User, error)
	DeleteSessionUser(ctx context.Context, session sessions.Session) error
	GetUserPassword(ctx context.Context, username string) (*string, error)
}
