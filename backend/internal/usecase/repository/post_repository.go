package repository

import (
	"context"
	"myapp/internal/entity"
)

type PostRepository interface {
	GetAll(context.Context) ([]entity.Post, error)
	Get(context.Context, int) (*entity.Post, error)
	CreatePost(context.Context, *entity.Post) (*entity.Post, error)
	UpdatePost(ctx context.Context, id int, title string, body string) (*entity.Post, error)
	DeletePost(ctx context.Context, id int) error
}
