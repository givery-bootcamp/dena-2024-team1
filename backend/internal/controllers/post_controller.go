package controllers

import (
	"errors"
	"myapp/internal/repositories"
	"myapp/internal/usecases"

	"github.com/gin-gonic/gin"
)

func GetPosts(ctx *gin.Context) {

	postRepository := repositories.NewPostRepository(DB(ctx))
	userRepository := repositories.NewUserRepository(DB(ctx))
	usecase := usecases.NewPostUsecase(userRepository, postRepository)
	result, err := usecase.GetPosts()
	if err != nil {
		handleError(ctx, 500, err)
	} else if result != nil {
		ctx.JSON(200, result)
	} else {
		handleError(ctx, 404, errors.New("not found"))
	}
}
