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
	su  usecase.SketchUsecase
	uu  usecase.UserUsecase
	ssu usecase.SSEUsecase
}

func NewSketchHandler(su usecase.SketchUsecase, uu usecase.UserUsecase, ssu usecase.SSEUsecase) SketchHandler {
	return SketchHandler{
		su:  su,
		uu:  uu,
		ssu: ssu,
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
	// userを取得
	session := sessions.Default(ctx)
	user, err := h.uu.GetSessionUser(ctx, session)

	if err != nil {
		handleError(ctx, 500, err)
		return
	}

	sketch, err := h.su.CreateSketch(ctx, &file, user.ID)
	if err != nil {
		handleError(ctx, 500, err)
	}

	// SSE
	h.ssu.NotifyClients("sketch created")

	ctx.JSON(201, openapi.CreateScketchesResponse{
		Id:        sketch.ID,
		ImageUrl:  config.S3BucketURL + sketch.ImageName,
		UserId:    sketch.UserID,
		UserName:  user.Name,
		CreatedAt: sketch.CreatedAt,
		UpdatedAt: sketch.UpdatedAt,
	})
}

func (h *SketchHandler) GetSketches(ctx *gin.Context) {
	ss, err := h.su.GetSketches(ctx)
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
