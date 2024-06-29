package sse

import (
	"log"
)

var broker *Broker = nil

func InitBroker() {
	if broker != nil {
		return
	}

	broker = newBroker()

	go broker.listen()
}

func GetBroker() *Broker {
	if broker == nil {
		return nil
	}

	return broker
}

type Broker struct {
	Clients      map[chan string]bool
	NewClients   chan chan string
	CloseClients chan chan string
	Messages     chan string
}

func newBroker() *Broker {
	return &Broker{
		Clients:      make(map[chan string]bool),
		NewClients:   make(chan chan string),
		CloseClients: make(chan chan string),
		Messages:     make(chan string),
	}
}

func (broker *Broker) listen() {
	for {
		select {
		case client := <-broker.NewClients:
			broker.Clients[client] = true
			log.Printf("Client added. %d registered clients", len(broker.Clients))
		case client := <-broker.CloseClients:
			delete(broker.Clients, client)
			close(client)
			log.Printf("Client removed. %d registered clients", len(broker.Clients))
		case msg := <-broker.Messages:
			for client := range broker.Clients {
				client <- msg
			}
		}
	}
}
