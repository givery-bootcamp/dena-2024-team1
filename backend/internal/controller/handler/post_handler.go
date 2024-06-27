package handler

import (
	"errors"
	"myapp/internal/entity"
	"myapp/internal/openapi"
	"myapp/internal/usecase"
	"strconv"

	"github.com/gin-contrib/sessions"
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
			Body:      p.Body,
			CreatedAt: p.CreatedAt,
			Id:        p.ID,
			Title:     p.Title,
			UpdatedAt: p.UpdatedAt,
			UserId:    p.UserID,
			UserName:  p.UserName,
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
			Body:      p.Body,
			CreatedAt: p.CreatedAt,
			Id:        p.ID,
			Title:     p.Title,
			UpdatedAt: p.UpdatedAt,
			UserId:    p.UserID,
			UserName:  p.UserName,
		}
		ctx.JSON(200, response)
	} else {
		handleError(ctx, 404, errors.New("not found"))
	}
}

func (h *PostHandler) CreatePost(ctx *gin.Context) {
	var request openapi.CreatePostRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		handleError(ctx, 400, err)
		return
	}
	session := sessions.Default(ctx)

	createdPost, err := h.pu.CreatePost(entity.Post{
		Title: request.Title,
		Body:  request.Body,
	}, session)
	if err != nil {
		handleError(ctx, 500, err)
		return
	}

	response := openapi.CreatePostResponse{
		Body:      createdPost.Body,
		CreatedAt: createdPost.CreatedAt,
		Id:        createdPost.ID,
		Title:     createdPost.Title,
		UpdatedAt: createdPost.UpdatedAt,
		UserId:    createdPost.UserID,
		UserName:  createdPost.UserName,
	}
	ctx.JSON(201, response)
}

func (h *PostHandler) UpdatePost(ctx *gin.Context) {
	// リクエストのbodyを取得、バインドをすることで他のものが入ってきたらエラーを返すようにする
	var request openapi.UpdatePostRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		handleError(ctx, 400, err)
		return
	}
	session := sessions.Default(ctx)

	// idを取得、エラーがあればエラーを返す
	// getPostからそのまま持ってきている
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handleError(ctx, 400, err)
	}

	// 代入したものでupdateする
	updatePost, err := h.pu.UpdatePost(id, request.Title, request.Body, session)
	if err != nil {
		handleError(ctx, 500, err)
		return
	}

	response := openapi.UpdatePostResponse{
		Body:      updatePost.Body,
		CreatedAt: updatePost.CreatedAt,
		Id:        updatePost.ID,
		Title:     updatePost.Title,
		UpdatedAt: updatePost.UpdatedAt,
		UserId:    updatePost.UserID,
		UserName:  updatePost.UserName,
	}
	ctx.JSON(200, response)
}

func (h *PostHandler) DeletePost(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handleError(ctx, 400, err)
	}
	session := sessions.Default(ctx)

	err = h.pu.DeletePost(id, session)
	if err != nil {
		handleError(ctx, 500, err)
	} else {
		ctx.JSON(204, nil)
	}
}
