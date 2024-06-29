package usecase

import "myapp/internal/utils/sse"

type SSEUsecase struct{}

func NewSSEUsecase() SSEUsecase {
	return SSEUsecase{}
}

func (u SSEUsecase) NotifyClients(message string) {
	broker := sse.GetBroker()
	if broker == nil {
		return
	}

	broker.Messages <- message
}
