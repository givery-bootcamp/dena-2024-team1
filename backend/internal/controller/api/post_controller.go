package api

import (
	"errors"
	"myapp/internal/controller/repository"
	"myapp/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPosts(ctx *gin.Context) {

	postRepository := repository.NewPostRepository(DB(ctx))
	usecase := usecase.NewPostUsecase(postRepository)
	result, err := usecase.GetPosts()
	if err != nil {
		handleError(ctx, 500, err)
	} else if result != nil {
		ctx.JSON(200, result)
	} else {
		handleError(ctx, 404, errors.New("not found"))
	}
}

func GetPost(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handleError(ctx, 400, err)
	}
	postRepository := repository.NewPostRepository(DB(ctx))
	usecase := usecase.NewPostUsecase(postRepository)
	result, err := usecase.GetPost(id)
	if err != nil {
		handleError(ctx, 500, err)
	} else if result != nil {
		ctx.JSON(200, result)
	} else {
		handleError(ctx, 404, errors.New("not found"))
	}
}
