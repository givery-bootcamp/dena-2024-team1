package repository

import (
	"errors"
	"myapp/internal/controller/repository/model"
	"myapp/internal/entity"

	repositoryIF "myapp/internal/usecase/repository"

	"gorm.io/gorm"
)

type HelloWorldRepository struct {
	Conn *gorm.DB
}

func NewHelloWorldRepository(conn *gorm.DB) repositoryIF.HelloWorldRepository {
	return &HelloWorldRepository{
		Conn: conn,
	}
}

func (r *HelloWorldRepository) Get(lang string) (*entity.HelloWorld, error) {
	var obj model.HelloWorld
	result := r.Conn.Where("lang = ?", lang).First(&obj)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return convertHelloWorldRepositoryModelToEntity(&obj), nil
}

func convertHelloWorldRepositoryModelToEntity(v *model.HelloWorld) *entity.HelloWorld {
	return &entity.HelloWorld{
		Lang:    v.Lang,
		Message: v.Message,
	}
}
