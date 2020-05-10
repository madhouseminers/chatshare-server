package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/madhouseminers/chatshare-server/internal/clients"
	"log"
	"math/rand"
	"strconv"
)

type handler struct {
	conn *websocket.Conn
	name *string
	bus  messageBus
}

func createHandler(conn *websocket.Conn, bus messageBus) *handler {
	var name = "WEBSOCKET" + strconv.Itoa(rand.Int())
	h := &handler{
		conn: conn,
		name: &name,
		bus:  bus,
	}

	go func() {
		h.bus.AddClient(h)
		h.startMessageLoop()
		h.bus.RemoveClient(h)
		err := h.conn.Close()
		if err != nil {
			log.Println("Unable to close handler: " + err.Error())
			return
		}
	}()

	return h
}

func (h *handler) startMessageLoop() {
	for {
		_, message, err := h.conn.ReadMessage()
		if err != nil {
			log.Println("Error in socket: " + err.Error())
			break
		}

		h.bus.Broadcast(clients.CreateMessage(string(message), h))
	}
}

func (h *handler) SendMessage(message *clients.Message) {
	err := h.conn.WriteMessage(1, []byte(message.GetContent()))
	if err != nil {
		log.Println("Unable to send message: " + err.Error())
		return
	}
}

func (h *handler) GetName() *string {
	return h.name
}
