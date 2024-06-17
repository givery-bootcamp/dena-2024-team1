package handler

import (
	"errors"
	"myapp/internal/entity"
	"myapp/internal/usecase"
	"strconv"
	"time"

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
	result, err := h.pu.GetPosts()
	if err != nil {
		handleError(ctx, 500, err)
	} else if result != nil {
		ctx.JSON(200, result)
	} else {
		handleError(ctx, 404, errors.New("not found"))
	}
}

func (h *PostHandler) GetPost(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handleError(ctx, 400, err)
	}
	result, err := h.pu.GetPost(id)
	if err != nil {
		handleError(ctx, 500, err)
	} else if result != nil {
		ctx.JSON(200, result)
	} else {
		handleError(ctx, 404, errors.New("not found"))
	}
}

// ここに追加して良いのか？？？自信ねぇ
type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

func (h *PostHandler) CreatePost(ctx *gin.Context) {
	var post usecase.PostUsecase
	if err := ctx.ShouldBindJSON(&post); err != nil {
		handleError(ctx, 400, err)
		return
	}

	createdPost, err := h.pu.CreatePost(entity.Post{})
	if err != nil {
		handleError(ctx, 500, err)
		return
	} else {
		ctx.JSON(201, createdPost)
	}
}
