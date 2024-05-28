package repositories

import (
	"errors"
	"fmt"
	"myapp/internal/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

// This struct is same as entity model
// However define again for training
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	gorm.Model
}

func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{
		Conn: conn,
	}
}

func (r *UserRepository) GetAll() ([]entities.User, error) {
	var users []User
	result := r.Conn.Find(users)
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

func convertUserRepositoryModelToEntity(ps []User) []entities.User {
	var users []entities.User

	for _, p := range ps {
		users = append(users, entities.User{
			Name:      p.Name,
			Password:  p.Password,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}
	return users
}
