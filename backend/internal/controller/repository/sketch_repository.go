package repository

import (
	"context"
	"fmt"
	"mime/multipart"

	"myapp/internal/entity"
	repositoryIF "myapp/internal/usecase/repository"

	"myapp/internal/controller/repository/ent"
	"myapp/internal/infrastructure/filestorage"

	"github.com/google/uuid"
)

type SketchRepository struct {
	Conn *ent.Client
}

func NewSketchRepository(conn *ent.Client) repositoryIF.SketchRepository {
	return &SketchRepository{
		Conn: conn,
	}
}

func (r *SketchRepository) CreateSketch(ctx context.Context, file *multipart.File) error {
	fn := uuid.New().String() + ".png"

	s3FileStorage := filestorage.SetUpS3()
	err := s3FileStorage.UploadFile(file, fn)
	if err != nil {
		return err
	}

	_, err = r.Conn.Sketch.
		Create().
		SetImageName(fn).
		SetUserID(1).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to create sketch: %w", err)
	}

	return nil
}

func (r *SketchRepository) GetAll(ctx context.Context) ([]entity.Sketch, error) {
	ss, err := r.Conn.Sketch.
		Query().
		WithUser().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all sketches: %w", err)
	}

	var sketches []entity.Sketch
	for _, s := range ss {
		sketches = append(sketches, entity.Sketch{
			ID:        s.ID,
			ImageName: s.ImageName,
			UserID:    s.UserID,
			UserName:  s.Edges.User.Name,
			CreatedAt: s.CreatedAt,
			UpdatedAt: s.UpdatedAt,
		})
	}

	return sketches, nil
}
