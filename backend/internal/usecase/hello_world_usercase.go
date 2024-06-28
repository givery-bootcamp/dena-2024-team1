package usecase

import (
	"context"
	"myapp/internal/entity"
	"myapp/internal/usecase/repository"
)

type HelloWorldUsecase struct {
	repository repository.HelloWorldRepository
}

func NewHelloWorldUsecase(r repository.HelloWorldRepository) HelloWorldUsecase {
	return HelloWorldUsecase{
		repository: r,
	}
}

func (u *HelloWorldUsecase) Execute(ctx context.Context, lang string) (*entity.HelloWorld, error) {
	return u.repository.Get(ctx, lang)
}
