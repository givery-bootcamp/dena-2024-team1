package usecase

import (
	"errors"
	"mime/multipart"
	"myapp/internal/entity"
	"myapp/internal/usecase/repository"

	"github.com/gin-contrib/sessions"
)

type SketchUsecase struct {
	sketchRepository repository.SketchRepository
}

func NewSketchUsecase(pr repository.SketchRepository) SketchUsecase {
	return SketchUsecase{
		sketchRepository: pr,
	}
}

func (u *SketchUsecase) CreateSketch(file *multipart.File, session sessions.Session) error {
	userID := session.Get("user_id").(int)
	if userID == 0 {
		return errors.New("user_id is empty")
	}

	return u.sketchRepository.CreateSketch(file, session)
}

func (u *SketchUsecase) GetSketches() ([]entity.Sketch, error) {
	return u.sketchRepository.GetAll()
}
