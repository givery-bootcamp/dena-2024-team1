package usecase

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type SSEUsecase struct {
	broker *Broker
}

func NewSSEUsecase() SSEUsecase {
	broker := NewBroker()
	go broker.Listen()

	return SSEUsecase{
		broker: broker,
	}
}

type Broker struct {
	clients      map[chan string]bool
	newClients   chan chan string
	closeClients chan chan string
	messages     chan string
}

func NewBroker() *Broker {
	return &Broker{
		clients:      make(map[chan string]bool),
		newClients:   make(chan chan string),
		closeClients: make(chan chan string),
		messages:     make(chan string),
	}
}

func (broker *Broker) Listen() {
	for {
		select {
		case client := <-broker.newClients:
			broker.clients[client] = true
			log.Printf("Client added. %d registered clients", len(broker.clients))
		case client := <-broker.closeClients:
			delete(broker.clients, client)
			close(client)
			log.Printf("Client removed. %d registered clients", len(broker.clients))
		case msg := <-broker.messages:
			for client := range broker.clients {
				client <- msg
			}
		}
	}
}

func (u SSEUsecase) ServeHTTP(ctx *gin.Context) {
	messageChan := make(chan string)

	broker := u.broker

	broker.newClients <- messageChan

	defer func() {
		broker.closeClients <- messageChan
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

func (u SSEUsecase) NotifyClients(message string) {
	broker := u.broker
	broker.messages <- message
}
