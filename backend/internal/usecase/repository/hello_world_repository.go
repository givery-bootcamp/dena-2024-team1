// usecase層とrepository層の間のinterface
package repository

import (
	"myapp/internal/entity"
)

type HelloWorldRepository interface {
	Get(lang string) (*entity.HelloWorld, error)
}
