package repository

import (
	"errors"
	"mime/multipart"

	"myapp/internal/controller/repository/model"
	"myapp/internal/entity"
	repositoryIF "myapp/internal/usecase/repository"

	"myapp/internal/infrastructure/filestorage"

	"github.com/google/uuid"
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

func (r *SketchRepository) CreateSketch(file *multipart.File, userID int) (entity.Sketch, error) {
	fn := uuid.New().String() + ".png"

	s3FileStorage := filestorage.SetUpS3()
	err := s3FileStorage.UploadFile(file, fn)
	if err != nil {
		return entity.Sketch{}, err
	}
	sketch := model.Sketch{
		ImageName: fn,
		UserID:    userID,
	}
	sketchResult := r.Conn.Create(&sketch)
	if sketchResult.Error != nil {
		return entity.Sketch{}, sketchResult.Error
	}
	return entity.Sketch{
		ID:        int(sketch.ID),
		ImageName: sketch.ImageName,
		UserID:    sketch.UserID,
		CreatedAt: sketch.CreatedAt,
		UpdatedAt: sketch.UpdatedAt,
	}, nil
}

func (r *SketchRepository) GetAll() ([]entity.Sketch, error) {
	var sketches []model.Sketch
	// データベースのエラー定義
	var ErrDatabase = errors.New("database error")
	sketchResult := r.Conn.Find(&sketches)
	if sketchResult.Error != nil {
		// レコードが見つからない場合のエラー定義
		if errors.Is(sketchResult.Error, gorm.ErrRecordNotFound) {
			return []entity.Sketch{}, nil
		}
		return nil, ErrDatabase
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
