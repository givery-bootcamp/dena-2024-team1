package handler

import (
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

type SignupRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h UserHandler) Signup(ctx *gin.Context) {
	var req SignupRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleError(ctx, 400, err)
		return
	}

	err := h.uu.Signup(req.Username, req.Password)

	if err != nil {
		handleError(ctx, 500, err)
	} else {
		ctx.JSON(200, gin.H{"message": "success"})
	}
}
