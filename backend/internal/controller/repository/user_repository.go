package repository

import (
	"encoding/json"
	"errors"
	"myapp/internal/config"
	"myapp/internal/controller/repository/model"
	"myapp/internal/entity"
	repositoryIF "myapp/internal/usecase/repository"

	"github.com/gin-contrib/sessions"
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

func NewUserRepository(conn *gorm.DB) repositoryIF.UserRepository {
	return &UserRepository{
		Conn: conn,
	}
}

func (r *UserRepository) GetAll() ([]entity.User, error) {
	var users []model.User
	result := r.Conn.Find(&users)
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
			ID:        int(p.ID),
			Name:      p.Name,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}
	return users
}

func (r *UserRepository) CreateUser(username, password string) (entity.User, error) {
	user := User{
		Name:     username,
		Password: password,
	}
	result := r.Conn.Create(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return entity.User{
		ID:        int(user.ID),
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (r *UserRepository) GetUserByUsername(username string) (entity.User, error) {
	var user User
	result := r.Conn.Where("name = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.User{}, nil
		}
		return entity.User{}, result.Error
	}
	return entity.User{
		ID:        int(user.ID),
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (r *UserRepository) GetUserPassword(username string) (string, error) {
	var user User
	result := r.Conn.Where("name = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "", nil
		}
		return "", result.Error
	}
	return user.Password, nil
}

// Sessionに保存するユーザー情報の構造体
type SesssionUser struct {
	ID   int
	Name string
}

func (r *UserRepository) SaveSession(session sessions.Session, user entity.User) error {
	// セッションに保存するユーザー情報
	sessionUser := SesssionUser{
		ID:   user.ID,
		Name: user.Name,
	}

	// セッションに保存するユーザー情報をJSONに変換
	sessionUserJson, err := json.Marshal(sessionUser)

	if err != nil {
		return err
	}

	// セッションに保存
	session.Set(config.SessionKey, string(sessionUserJson))

	err = session.Save()
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetSessionUser(session sessions.Session) (entity.User, error) {
	// セッションからユーザー情報を取得
	sessionUserJson, ok := session.Get(config.SessionKey).(string)
	if !ok {
		return entity.User{}, errors.New("session is empty")
	}

	// セッションから取得したユーザー情報を構造体に変換
	var sessionUser SesssionUser
	err := json.Unmarshal([]byte(sessionUserJson), &sessionUser)
	if err != nil {
		return entity.User{}, err
	}

	// ユーザー情報を取得
	user, err := r.GetUserByUsername(sessionUser.Name)

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *UserRepository) DeleteSessionUser(session sessions.Session) error {
	// セッションを削除
	session.Clear()
	err := session.Save()
	if err != nil {
		return err
	}
	return nil
}
