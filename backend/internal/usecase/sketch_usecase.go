package usecase

import (
	"myapp/internal/entity"
	"myapp/internal/usecase/repository"
)

type SketchUsecase struct {
	sketchRepository repository.SketchRepository
}

func NewSketchUsecase(pr repository.SketchRepository) SketchUsecase {
	return SketchUsecase{
		sketchRepository: pr,
	}
}

func (u *SketchUsecase) GetSketches() ([]entity.Sketch, error) {
	return u.sketchRepository.GetAll()
}
