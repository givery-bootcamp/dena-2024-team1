package handler

import (
	"myapp/internal/usecase"

	"github.com/gin-gonic/gin"
)

type SSEHandler struct {
	ssu usecase.SSEUsecase
}

func NewSSEHandler(ssu usecase.SSEUsecase) SSEHandler {
	return SSEHandler{
		ssu: ssu,
	}
}

func (h SSEHandler) ServeHTTP(ctx *gin.Context) {
	h.ssu.ServeHTTP(ctx)
}
