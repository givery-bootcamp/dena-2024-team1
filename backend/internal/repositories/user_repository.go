package repositories

import (
	"errors"
	"fmt"
	"myapp/internal/entities"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

// This struct is same as entity model
// However define again for training
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{
		Conn: conn,
	}
}

func (r *UserRepository) GetAll() ([]entities.Usr, error) {
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
			ID:        p.ID,
			Name:      p.Name,
			Password:  p.Password,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
			DeletedAt: p.DeletedAt,
		})
	}
	return users
}
