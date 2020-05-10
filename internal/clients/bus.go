package clients

import (
	"log"
)

type clientMessageBus struct {
	clients map[string]Client
}

func CreateBus() *clientMessageBus {
	return &clientMessageBus{clients: make(map[string]Client)}
}

func (b *clientMessageBus) AddClient(client Client) {
	log.Println("Added client: " + *client.GetName())
	b.clients[*client.GetName()] = client
}

func (b *clientMessageBus) RemoveClient(client Client) {
	log.Println("Removed client: " + *client.GetName())
	delete(b.clients, *client.GetName())
}

func (b *clientMessageBus) Broadcast(message *Message) {
	log.Println(message.GetContent())
	for _, client := range b.clients {
		if client.GetName() != message.GetSender().GetName() {
			client.SendMessage(message)
		}
	}
}
