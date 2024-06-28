package handler

import (
	"errors"
	"fmt"
	"myapp/internal/usecase"

	"github.com/gin-gonic/gin"
)

type HelloWorldHandler struct {
	hu usecase.HelloWorldUsecase
}

func NewHelloWorldHandler(hu usecase.HelloWorldUsecase) HelloWorldHandler {
	return HelloWorldHandler{
		hu: hu,
	}
}

func (h *HelloWorldHandler) HelloWorld(ctx *gin.Context) {
	lang := ctx.DefaultQuery("lang", "ja")
	if err := validateHelloWorldParameters(lang); err != nil {
		handleError(ctx, 400, err)
		return
	}
	result, err := h.hu.Execute(ctx, lang)
	if err != nil {
		handleError(ctx, 500, err)
	} else if result != nil {
		ctx.JSON(200, result)
	} else {
		handleError(ctx, 404, errors.New("not found"))
	}
}

func validateHelloWorldParameters(lang string) error {
	if len(lang) != 2 {
		return fmt.Errorf("invalid lang parameter: %s", lang)
	}
	return nil
}
