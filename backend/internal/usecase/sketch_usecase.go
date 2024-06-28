package usecase

import (
	"context"
	"mime/multipart"
	"myapp/internal/entity"
	"myapp/internal/usecase/repository"
)

type SketchUsecase struct {
	sketchRepository repository.SketchRepository
}

func NewSketchUsecase(pr repository.SketchRepository, ur repository.UserRepository) SketchUsecase {
	return SketchUsecase{
		sketchRepository: pr,
	}
}

func (u *SketchUsecase) CreateSketch(ctx context.Context, file *multipart.File, userid int) (*entity.Sketch, error) {
	return u.sketchRepository.CreateSketch(ctx, file, userid)
}

func (u *SketchUsecase) GetSketches(ctx context.Context) ([]entity.Sketch, error) {
	return u.sketchRepository.GetAll(ctx)
}

func (u *SketchUsecase) DeleteSketch(ctx context.Context, id int) error {
	return u.sketchRepository.DeleteSketch(ctx, id)
}
