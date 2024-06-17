package repository

import (
	"errors"
	"myapp/internal/entity"
	repositoryIF "myapp/internal/usecase/repository"

	"gorm.io/gorm"
)

type PostRepository struct {
	Conn *gorm.DB
}

// This struct is same as entity model
// However define again for training
type Post struct {
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`

	gorm.Model
}

func NewPostRepository(conn *gorm.DB) repositoryIF.PostRepository {
	return &PostRepository{
		Conn: conn,
	}
}

func (r *PostRepository) CreatePost(post *entity.Post) (entity.Post, error) {
	postResult := r.Conn.Create(post)
	if postResult.Error != nil {
		return entity.Post{}, postResult.Error
	}
	return *post, nil
}

func (r *PostRepository) GetAll() ([]entity.Post, error) {
	var posts []Post
	postResult := r.Conn.Find(&posts)
	if postResult.Error != nil {
		if errors.Is(postResult.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, postResult.Error
	}
	var users []User
	userResult := r.Conn.Find(&users)
	if userResult.Error != nil {
		if errors.Is(userResult.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, userResult.Error
	}
	return convertPostsRepositoryModelToEntities(posts, users), nil
}

func convertPostsRepositoryModelToEntities(ps []Post, us []User) []entity.Post {
	var posts []entity.Post

	for _, p := range ps {
		post := entity.Post{
			ID:        int(p.ID),
			Title:     p.Title,
			Body:      p.Body,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		}
		for _, u := range us {
			if p.UserID == int(u.ID) {
				post.UserName = u.Name
				break
			}
		}
		posts = append(posts, post)
	}
	return posts
}

func (r *PostRepository) Get(id int) (*entity.Post, error) {
	var post Post
	postResult := r.Conn.Where("id = ?", id).First(&post)
	if postResult.Error != nil {
		if errors.Is(postResult.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, postResult.Error
	}
	var user User
	userResult := r.Conn.Where("id = ?", post.UserID).First(&user)
	if userResult.Error != nil {
		if errors.Is(userResult.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, userResult.Error
	}

	return convertPostRepositoryModelToEntity(&post, &user), nil
}

func convertPostRepositoryModelToEntity(p *Post, u *User) *entity.Post {
	return &entity.Post{
		ID:        int(p.ID),
		Title:     p.Title,
		Body:      p.Body,
		UserName:  u.Name,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
