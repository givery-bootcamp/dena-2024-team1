package handler

import (
	"errors"
	"myapp/internal/config"
	"myapp/internal/openapi"
	"myapp/internal/usecase"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SketchHandler struct {
	su usecase.SketchUsecase
	uu usecase.UserUsecase
}

func NewSketchHandler(su usecase.SketchUsecase) SketchHandler {
	return SketchHandler{
		su: su,
	}
}

func (h *SketchHandler) CreateSketch(ctx *gin.Context) {
	const MaxUploadSize = 10 * 1024 * 1024 // 10MB

	requestFile, err := ctx.FormFile("file")
	if err != nil {
		handleError(ctx, 400, err)
		return
	}

	// Check file size
	if requestFile.Size > MaxUploadSize {
		handleError(ctx, 400, errors.New("file size too large"))
		return
	}
	// Check file type
	if requestFile.Header.Get("Content-Type") != "image/png" && requestFile.Header.Get("Content-Type") != "image/jpeg" {
		handleError(ctx, 400, errors.New("file type not allowed"))
		return
	}

	file, err := requestFile.Open()
	if err != nil {
		handleError(ctx, 500, err)
		return
	}
	session := sessions.Default(ctx)
	user, err := h.uu.GetSessionUser(session)

	if err != nil {
		handleError(ctx, 500, err)
		return
	} else {
		ctx.JSON(200, gin.H{"user_id": user.ID})
	}

	err = h.su.CreateSketch(&file, session)
	if err != nil {
		handleError(ctx, 500, err)
	} else {
		ctx.JSON(201, nil)
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
