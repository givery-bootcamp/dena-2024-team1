package repository

import (
	"context"
	"mime/multipart"
	"myapp/internal/entity"
)

type SketchRepository interface {
	CreateSketch(ctx context.Context, file *multipart.File) error
	GetAll(ctx context.Context) ([]entity.Sketch, error)
}
