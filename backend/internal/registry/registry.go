// DIコンテナ。各層の依存関係を解決する。
package registry

import (
	"myapp/internal/controller/handler"
	repoImp "myapp/internal/controller/repository"
	"myapp/internal/infrastructure/database"
	"myapp/internal/infrastructure/filestorage"
	"myapp/internal/usecase"
	repoIf "myapp/internal/usecase/repository"
)

type APIHandler struct {
	HelloWorldHandler handler.HelloWorldHandler
	PostHandler       handler.PostHandler
	UserHandler       handler.UserHandler
	SketchHandler     handler.SketchHandler
	SSEHandler        handler.SSEHandler
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
	uh := handler.NewUserHandler(u.uh)
	sh := handler.NewSketchHandler(u.sh, u.uh, u.ssu)
	ssh := handler.NewSSEHandler(u.ssu)

	return &APIHandler{
		HelloWorldHandler: hh,
		PostHandler:       ph,
		UserHandler:       uh,
		SketchHandler:     sh,
		SSEHandler:        ssh,
	}
}

type apiUsecase struct {
	hu  usecase.HelloWorldUsecase
	ph  usecase.PostUsecase
	uh  usecase.UserUsecase
	sh  usecase.SketchUsecase
	ssu usecase.SSEUsecase
}

func newAPIUsecase(r apiRepository) *apiUsecase {
	hu := usecase.NewHelloWorldUsecase(r.hr)
	ph := usecase.NewPostUsecase(r.pr, r.ur)
	uh := usecase.NewUserUsecase(r.ur)
	sh := usecase.NewSketchUsecase(r.sh, r.ur)
	ssu := usecase.NewSSEUsecase()

	return &apiUsecase{
		hu:  hu,
		ph:  ph,
		uh:  uh,
		sh:  sh,
		ssu: ssu,
	}
}

type apiRepository struct {
	hr repoIf.HelloWorldRepository
	pr repoIf.PostRepository
	ur repoIf.UserRepository
	sh repoIf.SketchRepository
}

func newAPIRepository() *apiRepository {
	db := database.SetupDB()
	str := filestorage.SetUpS3()

	hr := repoImp.NewHelloWorldRepository(db)
	pr := repoImp.NewPostRepository(db)
	ur := repoImp.NewUserRepository(db)
	sh := repoImp.NewSketchRepository(db, str)

	return &apiRepository{
		hr: hr,
		pr: pr,
		ur: ur,
		sh: sh,
	}
}
