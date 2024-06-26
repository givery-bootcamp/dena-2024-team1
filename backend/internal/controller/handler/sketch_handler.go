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

func (h *SketchHandler) CreateSketch(ctx *gin.Context) {
	var request openapi.CreateScketchesRequest
	const MaxUploadSize = 10 * 1024 * 1024 // 10MB

	if err := ctx.ShouldBindJSON(&request); err != nil {
		handleError(ctx, 400, err)
	}

	// file, err := ctx.FormFile("file")
	// if err != nil {
	// 	handleError(ctx, 400, err)
	// 	return
	// }
	// // Check file size
	// if file.Size > MaxUploadSize {
	// 	handleError(ctx, 400, errors.New("file size too large"))
	// 	return
	// }
	// // Check file type
	// if file.Header.Get("Content-Type") != "image/png" && file.Header.Get("Content-Type") != "image/jpeg" {
	// 	handleError(ctx, 400, errors.New("file type not allowed"))
	// 	return
	// }

	// // Save file
	// // 保存先が分かってないかもしれない
	// // destination := "./uploads/" + file.Filename

	// if err := ctx.SaveUploadedFile(file, destination); err != nil {
	// 	ctx.String(http.StatusInternalServerError, "Failed to save file: %s", err.Error())
	// 	return
	// }

	err := h.su.CreateSketch(*request.Filename, request.File)
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
