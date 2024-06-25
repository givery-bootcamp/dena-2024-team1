package handler

import (
	"errors"
	"myapp/internal/config"
	"myapp/internal/openapi"
	"myapp/internal/usecase"

	"github.com/gin-gonic/gin"
)

type SketchHandler struct {
	su usecase.SketchUsecase
}

func NewSketchHandler(su usecase.SketchUsecase) SketchHandler {
	return SketchHandler{
		su: su,
	}
}

func (h *SketchHandler) GetSketches(ctx *gin.Context) {
	ss, err := h.su.GetSketches()
	var response openapi.GetAllSketchesResponse
	for _, s := range ss {
		imageURL := config.S3BucketURL + s.ImageName
		response = append(response, openapi.Sketch{
			Id:        s.ID,
			ImageUrl:  imageURL,
			UserId:    s.UserID,
			UserName:  s.UserName,
			CreatedAt: s.CreatedAt,
			UpdatedAt: s.UpdatedAt,
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
