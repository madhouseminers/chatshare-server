package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/madhouseminers/chatshare-server/internal/clients"
	"log"
	"os"
	"strings"
)

type handler struct {
	conn *websocket.Conn
	name *string
	bus  messageBus
}

func createHandler(conn *websocket.Conn, bus messageBus) *handler {
	h := &handler{
		conn: conn,
		bus:  bus,
	}

	go func() {
		err := h.conn.WriteMessage(1, []byte("HELLO"))
		if err != nil {
			err = h.conn.Close()
			return
		}
		h.startMessageLoop()
		if h.name != nil {
			h.bus.Direct(clients.CreateMessage(*h.name+" has disconnected", h), "Discord")
			h.bus.RemoveClient(h)
		}
		err = h.conn.Close()
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

		if h.name == nil {
			auth := strings.SplitN(string(message), "::", 2)
			log.Println("Got authentication message from: " + auth[0])
			if auth[1] != os.Getenv("chatsharePSK") {
				log.Println("Authentication failed from: " + auth[0])
				return
			}
			h.name = &auth[0]
			h.bus.AddClient(h)
			h.bus.Direct(clients.CreateMessage(*h.name+" has connected", h), "Discord")
		} else {
			log.Println("Got message: " + string(message))
			h.bus.Broadcast(clients.CreateMessage(string(message), h))
		}
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
