package repository

import (
	"errors"
	"fmt"
	"myapp/internal/controller/repository/model"
	"myapp/internal/entity"
	repositoryIF "myapp/internal/usecase/repository"

	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) repositoryIF.UserRepository {
	return &UserRepository{
		Conn: conn,
	}
}

func (r *UserRepository) GetAll() ([]entity.User, error) {
	var users []model.User
	result := r.Conn.Find(&users)
	fmt.Printf("%+v\n", result)
	fmt.Printf("%+v\n", users)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return convertUserRepositoryModelToEntity(users), nil
}

func convertUserRepositoryModelToEntity(ps []model.User) []entity.User {
	var users []entity.User

	for _, p := range ps {
		users = append(users, entity.User{
			Name:      p.Name,
			Password:  p.Password,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}
	return users
}
