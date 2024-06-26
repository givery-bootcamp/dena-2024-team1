package repository

import (
	"errors"
	// "go/types"
	"myapp/internal/controller/repository/model"
	"myapp/internal/entity"
	repositoryIF "myapp/internal/usecase/repository"

	"github.com/oapi-codegen/runtime/types"
	"gorm.io/gorm"
)

type SketchRepository struct {
	Conn *gorm.DB
}

func NewSketchRepository(conn *gorm.DB) repositoryIF.SketchRepository {
	return &SketchRepository{
		Conn: conn,
	}
}

func (r *SketchRepository) CreateSketch(filename string, file *types.File) (entity.Sketch, error) {
	// types.Fileを使ってるけど、関数内で使っていないので、削除しても問題ないのかな？
	// そもそもRepositoryの責務ってなんだ？ってなっている
	sketch := model.Sketch{
		ImageName: filename,
	}
	sketchResult := r.Conn.Create(&sketch)
	if sketchResult.Error != nil {
		return entity.Sketch{}, sketchResult.Error
	}
	return entity.Sketch{
		ID:        int(sketch.ID),
		ImageName: sketch.ImageName,
		CreatedAt: sketch.CreatedAt,
		UpdatedAt: sketch.UpdatedAt,
	}, nil
}

func (r *SketchRepository) GetAll() ([]entity.Sketch, error) {
	var sketches []model.Sketch
	sketchResult := r.Conn.Find(&sketches)
	if sketchResult.Error != nil {
		if errors.Is(sketchResult.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, sketchResult.Error
	}
	var users []model.User
	userResult := r.Conn.Find(&users)
	if userResult.Error != nil {
		if errors.Is(userResult.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, userResult.Error
	}
	return convertSketchesRepositoryModelToEntities(sketches, users), nil
}

func convertSketchesRepositoryModelToEntities(ss []model.Sketch, us []model.User) []entity.Sketch {
	var sketches []entity.Sketch

	for _, s := range ss {
		sketch := entity.Sketch{
			ID:        int(s.ID),
			ImageName: s.ImageName,
			UserID:    s.UserID,
			CreatedAt: s.CreatedAt,
			UpdatedAt: s.UpdatedAt,
		}
		for _, u := range us {
			if s.UserID == int(u.ID) {
				sketch.UserName = u.Name
				break
			}
		}
		sketches = append(sketches, sketch)
	}
	return sketches
}
