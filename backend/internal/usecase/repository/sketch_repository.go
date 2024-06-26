package repository

import (
	"myapp/internal/entity"

	"github.com/oapi-codegen/runtime/types"
)

type SketchRepository interface {
	CreateSketch(filenanme string, file *types.File) error
	GetAll() ([]entity.Sketch, error)
}
