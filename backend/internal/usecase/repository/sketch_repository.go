package repository

import (
	"bytes"
	"myapp/internal/entity"
)

type SketchRepository interface {
	CreateSketch(reader *bytes.Reader) error
	GetAll() ([]entity.Sketch, error)
}
