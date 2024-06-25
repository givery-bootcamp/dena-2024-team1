package repository

import (
	"myapp/internal/entity"
)

type SketchRepository interface {
	GetAll() ([]entity.Sketch, error)
}
