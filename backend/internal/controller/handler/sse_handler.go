package handler

import (
	"fmt"
	"myapp/internal/usecase"
	"myapp/internal/utils/sse"

	"github.com/gin-gonic/gin"
)

type SSEHandler struct {
	ssu usecase.SSEUsecase
}

func NewSSEHandler(ssu usecase.SSEUsecase) SSEHandler {
	sse.InitBroker()

	return SSEHandler{
		ssu: ssu,
	}
}

func (h SSEHandler) ServeHTTP(ctx *gin.Context) {
	messageChan := make(chan string)

	broker := sse.GetBroker()

	if broker == nil {
		return
	}

	broker.NewClients <- messageChan

	defer func() {
		broker.CloseClients <- messageChan
	}()

	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")

	for {
		select {
		case msg := <-messageChan:
			fmt.Fprintf(ctx.Writer, "data: %s\n\n", msg)
			ctx.Writer.Flush()
		case <-ctx.Writer.CloseNotify():
			return
		}
	}
}
