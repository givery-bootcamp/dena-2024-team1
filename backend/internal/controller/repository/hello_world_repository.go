package repository

import (
	"context"
	"fmt"
	"myapp/internal/entity"
	"myapp/internal/infrastructure/database/ent"
	helloworldEntity "myapp/internal/infrastructure/database/ent/helloworld"

	repositoryIF "myapp/internal/usecase/repository"
)

type HelloWorldRepository struct {
	Conn *ent.Client
}

func NewHelloWorldRepository(conn *ent.Client) repositoryIF.HelloWorldRepository {
	return &HelloWorldRepository{
		Conn: conn,
	}
}

func (r *HelloWorldRepository) Get(ctx context.Context, lang string) (*entity.HelloWorld, error) {
	hw, err := r.Conn.HelloWorld.Query().Where(helloworldEntity.Lang(lang)).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get hello world message: %w", err)
	}

	return convertHelloWorldRepositoryModelToEntity(hw), nil
}

func convertHelloWorldRepositoryModelToEntity(v *ent.HelloWorld) *entity.HelloWorld {
	return &entity.HelloWorld{
		Lang:    v.Lang,
		Message: v.Message,
	}
}
