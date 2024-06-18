package repository

import (
	"errors"
	"myapp/internal/controller/repository/model"
	"myapp/internal/entity"
	repositoryIF "myapp/internal/usecase/repository"

	"gorm.io/gorm"
)

type PostRepository struct {
	Conn *gorm.DB
}

func NewPostRepository(conn *gorm.DB) repositoryIF.PostRepository {
	return &PostRepository{
		Conn: conn,
	}
}

func (r *PostRepository) CreatePost(post *entity.Post) (entity.Post, error) {
	// Convert entity.Post to repository.Post
	doPost := Post{
		UserID: post.UserID,
		Title:  post.Title,
		Body:   post.Body,
		Model: gorm.Model{
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		},
	}

	postResult := r.Conn.Create(&doPost)
	if postResult.Error != nil {
		return entity.Post{}, postResult.Error
	}

	resultPost := entity.Post{
		ID:        int(doPost.ID),
		Title:     doPost.Title,
		Body:      doPost.Body,
		UserID:    doPost.UserID,
		CreatedAt: doPost.CreatedAt,
		UpdatedAt: doPost.UpdatedAt,
	}
	return resultPost, nil
}

func (r *PostRepository) GetAll() ([]entity.Post, error) {
	var posts []model.Post
	postResult := r.Conn.Find(&posts)
	if postResult.Error != nil {
		if errors.Is(postResult.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, postResult.Error
	}
	var users []model.User
	userResult := r.Conn.Find(&users)
	if userResult.Error != nil {
		if errors.Is(userResult.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, userResult.Error
	}
	return convertPostsRepositoryModelToEntities(posts, users), nil
}

func convertPostsRepositoryModelToEntities(ps []model.Post, us []model.User) []entity.Post {
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
	var post model.Post
	postResult := r.Conn.Where("id = ?", id).First(&post)
	if postResult.Error != nil {
		if errors.Is(postResult.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, postResult.Error
	}
	var user model.User
	userResult := r.Conn.Where("id = ?", post.UserID).First(&user)
	if userResult.Error != nil {
		if errors.Is(userResult.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, userResult.Error
	}

	return convertPostRepositoryModelToEntity(&post, &user), nil
}

func convertPostRepositoryModelToEntity(p *model.Post, u *model.User) *entity.Post {
	return &entity.Post{
		ID:        int(p.ID),
		Title:     p.Title,
		Body:      p.Body,
		UserName:  u.Name,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
