package repository

import (
	"context"
	"fmt"
	"myapp/internal/controller/repository/ent"
	postEntity "myapp/internal/controller/repository/ent/post"
	"myapp/internal/entity"
	repositoryIF "myapp/internal/usecase/repository"
)

type PostRepository struct {
	Conn *ent.Client
}

func NewPostRepository(conn *ent.Client) repositoryIF.PostRepository {
	return &PostRepository{
		Conn: conn,
	}
}

func (r *PostRepository) CreatePost(ctx context.Context, p *entity.Post) (*entity.Post, error) {
	post, err := r.Conn.Post.
		Create().
		SetTitle(p.Title).
		SetBody(p.Body).
		SetUserID(p.UserID).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	resultPost := entity.Post{
		ID:        post.ID,
		Title:     post.Title,
		Body:      post.Body,
		UserID:    post.UserID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
	return &resultPost, nil
}

func (r *PostRepository) UpdatePost(ctx context.Context, id int, title string, body string) (*entity.Post, error) {
	post, err := r.Conn.Post.
		UpdateOneID(id).
		SetTitle(title).
		SetBody(body).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update post: %w", err)
	}

	resultPost := entity.Post{
		ID:        post.ID,
		Title:     post.Title,
		Body:      post.Body,
		UserID:    post.UserID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
	return &resultPost, nil
}

func (r *PostRepository) GetAll(ctx context.Context) ([]entity.Post, error) {
	ps, err := r.Conn.Post.
		Query().
		WithUser().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get posts: %w", err)
	}

	var posts []entity.Post
	for _, p := range ps {
		post := entity.Post{
			ID:        p.ID,
			Title:     p.Title,
			Body:      p.Body,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
			UserID:    p.Edges.User.ID,
			UserName:  p.Edges.User.Name,
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostRepository) Get(ctx context.Context, id int) (*entity.Post, error) {
	p, err := r.Conn.Post.
		Query().
		Where(postEntity.IDEQ(id)).
		WithUser().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get post: %w", err)
	}

	post := entity.Post{
		ID:        p.ID,
		Title:     p.Title,
		Body:      p.Body,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		UserID:    p.Edges.User.ID,
		UserName:  p.Edges.User.Name,
	}

	return &post, nil
}

func (r *PostRepository) DeletePost(ctx context.Context, id int) error {
	err := r.Conn.Post.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}

	return nil
}
