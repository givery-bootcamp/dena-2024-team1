// DIコンテナ。各層の依存関係を解決する。
package registry

import (
	"myapp/internal/controller/handler"
	repoImp "myapp/internal/controller/repository"
	"myapp/internal/infrastructure/database"
	"myapp/internal/usecase"
	repoIf "myapp/internal/usecase/repository"
)

type APIHandler struct {
	HelloWorldHandler handler.HelloWorldHandler
	PostHandler       handler.PostHandler
}

func NewAPIHandler() *APIHandler {
	r := newAPIRepository()
	u := newAPIUsecase(*r)
	h := newAPIHandler(*u)
	return h
}

func newAPIHandler(u apiUsecase) *APIHandler {
	hh := handler.NewHelloWorldHandler(u.hu)
	ph := handler.NewPostHandler(u.ph)

	return &APIHandler{
		HelloWorldHandler: hh,
		PostHandler:       ph,
	}
}

type apiUsecase struct {
	hu usecase.HelloWorldUsecase
	ph usecase.PostUsecase
}

func newAPIUsecase(r apiRepository) *apiUsecase {
	hu := usecase.NewHelloWorldUsecase(r.hr)
	ph := usecase.NewPostUsecase(r.pr)

	return &apiUsecase{
		hu: hu,
		ph: ph,
	}
}

type apiRepository struct {
	hr repoIf.HelloWorldRepository
	pr repoIf.PostRepository
	ur repoIf.UserRepository
}

func newAPIRepository() *apiRepository {
	db := database.SetupDB()

	hr := repoImp.NewHelloWorldRepository(db)
	pr := repoImp.NewPostRepository(db)
	ur := repoImp.NewUserRepository(db)

	return &apiRepository{
		hr: hr,
		pr: pr,
		ur: ur,
	}
}
