package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"myapp/internal/config"
	"myapp/internal/entity"
	"myapp/internal/infrastructure/database/ent"
	userEntity "myapp/internal/infrastructure/database/ent/user"
	repositoryIF "myapp/internal/usecase/repository"

	"github.com/gin-contrib/sessions"
	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *ent.Client
}

// This struct is same as entity model
// However define again for training
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	gorm.Model
}

func NewUserRepository(conn *ent.Client) repositoryIF.UserRepository {
	return &UserRepository{
		Conn: conn,
	}
}

func (r *UserRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	us, err := r.Conn.User.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}

	var users []entity.User
	for _, u := range us {
		users = append(users, entity.User{
			ID:        u.ID,
			Name:      u.Name,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}
	return users, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, username, password string) (*entity.User, error) {
	u, err := r.Conn.User.
		Create().
		SetName(username).
		SetPassword(password).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	user := entity.User{
		ID:        u.ID,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	return &user, nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	fmt.Println("get user by username")
	u, err := r.Conn.User.Query().Where(userEntity.Name(username)).First(ctx)
	fmt.Println(err)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	user := entity.User{
		ID:        u.ID,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	return &user, nil
}

func (r *UserRepository) GetUserPassword(ctx context.Context, username string) (*string, error) {
	u, err := r.Conn.User.Query().Where(userEntity.Name(username)).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &u.Password, nil
}

// Sessionに保存するユーザー情報の構造体
type SesssionUser struct {
	ID   int
	Name string
}

func (r *UserRepository) SaveSession(ctx context.Context, session sessions.Session, user entity.User) error {
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

func (r *UserRepository) GetSessionUser(ctx context.Context, session sessions.Session) (*entity.User, error) {
	// セッションからユーザー情報を取得
	sessionUserJson, ok := session.Get(config.SessionKey).(string)
	if !ok {
		return nil, errors.New("session is empty")
	}

	// セッションから取得したユーザー情報を構造体に変換
	var sessionUser SesssionUser
	err := json.Unmarshal([]byte(sessionUserJson), &sessionUser)
	if err != nil {
		return nil, err
	}

	// ユーザー情報を取得
	user, err := r.GetUserByUsername(ctx, sessionUser.Name)
	fmt.Printf("get session user: %+v\n", user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) DeleteSessionUser(ctx context.Context, session sessions.Session) error {
	// セッションを削除
	session.Clear()
	err := session.Save()
	if err != nil {
		return err
	}
	return nil
}
