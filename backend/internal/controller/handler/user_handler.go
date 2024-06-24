package handler

import (
	"myapp/internal/openapi"
	"myapp/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uu usecase.UserUsecase
}

func NewUserHandler(uu usecase.UserUsecase) UserHandler {
	return UserHandler{
		uu: uu,
	}
}

func (h UserHandler) Signup(ctx *gin.Context) {
	var req openapi.SignUpJSONBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleError(ctx, 400, err)
		return
	}

	err := h.uu.Signup(*req.Username, *req.Password)

	if err != nil {
		handleError(ctx, 500, err)
	} else {
		ctx.JSON(200, gin.H{"message": "success"})
	}
}

func (h UserHandler) Signin(ctx *gin.Context) {
	var req openapi.SignInJSONBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleError(ctx, 400, err)
		return
	}

	err := h.uu.Signin(*req.Username, *req.Password, ctx.Request, ctx.Writer)

	if err != nil {
		handleError(ctx, 500, err)
	} else {
		ctx.JSON(200, gin.H{"message": "success"})
	}
}

func (h UserHandler) GetSessionUser(ctx *gin.Context) {
	user, err := h.uu.GetSessionUser(ctx.Request)

	if err != nil {
		handleError(ctx, 500, err)
	} else {
		ctx.JSON(200, gin.H{"user": user})
	}
}

func (h UserHandler) Signout(ctx *gin.Context) {
	err := h.uu.Signout(ctx.Request, ctx.Writer)

	if err != nil {
		handleError(ctx, 500, err)
	} else {
		ctx.JSON(200, gin.H{"message": "success"})
	}
}
