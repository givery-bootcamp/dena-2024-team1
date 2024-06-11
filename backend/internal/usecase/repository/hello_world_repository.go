package repository

import (
	"myapp/internal/entity"
)

type HelloWorldRepository interface {
	Get(lang string) (*entity.HelloWorld, error)
}
