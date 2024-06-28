// usecase層とrepository層の間のinterface
package repository

import (
	"context"
	"myapp/internal/entity"
)

type HelloWorldRepository interface {
	Get(ctx context.Context, lang string) (*entity.HelloWorld, error)
}
