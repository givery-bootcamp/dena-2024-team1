package repository

import (
	"context"
	"mime/multipart"
	"myapp/internal/entity"
)

type SketchRepository interface {
	CreateSketch(ctx context.Context, file *multipart.File, userID int) (*entity.Sketch, error)
	GetAll(ctx context.Context) ([]entity.Sketch, error)
	DeleteSketch(ctx context.Context, id int) error
}
