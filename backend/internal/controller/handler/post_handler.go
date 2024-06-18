package handler

import (
	"errors"
	"myapp/internal/openapi"
	"myapp/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	pu usecase.PostUsecase
}

func NewPostHandler(pu usecase.PostUsecase) PostHandler {
	return PostHandler{
		pu: pu,
	}
}

func (h *PostHandler) GetPosts(ctx *gin.Context) {
	ps, err := h.pu.GetPosts()
	var response openapi.GetAllPostsResponse
	for _, p := range ps {
		response = append(response, openapi.Post{
			Body:      &p.Body,
			CreatedAt: &p.CreatedAt,
			Id:        &p.ID,
			Title:     &p.Title,
			UpdatedAt: &p.UpdatedAt,
			UserName:  &p.UserName,
		})
	}

	if err != nil {
		handleError(ctx, 500, err)
	} else if response != nil {
		ctx.JSON(200, response)
	} else {
		handleError(ctx, 404, errors.New("not found"))
	}
}

func (h *PostHandler) GetPost(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handleError(ctx, 400, err)
	}

	p, err := h.pu.GetPost(id)
	if err != nil {
		handleError(ctx, 500, err)
	}

	if p != nil {
		response := openapi.GetPostResponse{
			Body:      &p.Body,
			CreatedAt: &p.CreatedAt,
			Id:        &p.ID,
			Title:     &p.Title,
			UpdatedAt: &p.UpdatedAt,
			UserName:  &p.UserName,
		}
		ctx.JSON(200, response)
	} else {
		handleError(ctx, 404, errors.New("not found"))
	}
}
