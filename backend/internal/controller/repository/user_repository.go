package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"myapp/internal/config"
	"myapp/internal/entity"
	repositoryIF "myapp/internal/usecase/repository"
	"net/http"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

type UserRepository struct {
	Conn         *gorm.DB
	SessionStore *sessions.CookieStore
}

// This struct is same as entity model
// However define again for training
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	gorm.Model
}

func NewUserRepository(conn *gorm.DB, sessionStore *sessions.CookieStore) repositoryIF.UserRepository {
	return &UserRepository{
		Conn:         conn,
		SessionStore: sessionStore,
	}
}

func (r *UserRepository) GetAll() ([]entity.User, error) {
	var users []User
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

func convertUserRepositoryModelToEntity(ps []User) []entity.User {
	var users []entity.User

	for _, p := range ps {
		users = append(users, entity.User{
			Name:      p.Name,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}
	return users
}

func (r *UserRepository) CreateUser(username, password string) error {
	user := User{
		Name:     username,
		Password: password,
	}
	result := r.Conn.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
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

func (r *UserRepository) SaveSession(req *http.Request, w http.ResponseWriter, user entity.User) error {
	session, err := r.SessionStore.Get(req, config.SessionName)
	if err != nil {
		return err
	}

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
	session.Values[config.SessionKey] = string(sessionUserJson)
	fmt.Printf("save session: %+v\n", session.Values[config.SessionKey])

	err = session.Save(req, w)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetSessionUser(req *http.Request) (entity.User, error) {
	session, err := r.SessionStore.Get(req, config.SessionName)
	if err != nil {
		return entity.User{}, err
	}

	// セッションからユーザー情報を取得
	sessionUserJson, ok := session.Values[config.SessionKey].(string)
	if !ok {
		return entity.User{}, errors.New("session is empty")
	}

	// セッションから取得したユーザー情報を構造体に変換
	var sessionUser SesssionUser
	err = json.Unmarshal([]byte(sessionUserJson), &sessionUser)
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

func (r *UserRepository) DeleteSessionUser(req *http.Request, w http.ResponseWriter) error {
	session, err := r.SessionStore.Get(req, config.SessionName)
	if err != nil {
		return err
	}

	// セッションを削除
	session.Options.MaxAge = -1
	err = session.Save(req, w)
	if err != nil {
		return err
	}
	return nil
}
