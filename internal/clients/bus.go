package clients

import "fmt"

type clientMessageBus struct {
	clients map[string]Client
}

func CreateBus() *clientMessageBus {
	return &clientMessageBus{clients: make(map[string]Client)}
}

func (b *clientMessageBus) AddClient(client Client) {
	b.clients[*client.GetName()] = client
	fmt.Println(b.clients)
}

func (b *clientMessageBus) RemoveClient(client Client) {
	delete(b.clients, *client.GetName())
}

func (b *clientMessageBus) Broadcast(message *Message) {
	for _, client := range b.clients {
		if client.GetName() != message.GetSender().GetName() {
			client.SendMessage(message)
		}
	}
}
