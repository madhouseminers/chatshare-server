package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/madhouseminers/chatshare-server/internal/clients"
	"log"
	"os"
)

type handler struct {
	conn    *websocket.Conn
	name    *string
	bus     messageBus
	version string
}

func createHandler(conn *websocket.Conn, bus messageBus) *handler {
	h := &handler{
		conn:    conn,
		bus:     bus,
		version: "2.0",
	}

	go func() {
		err := h.conn.WriteMessage(1, []byte("HELLO"))
		if err != nil {
			err = h.conn.Close()
			return
		}
		h.startMessageLoop()
		if h.name != nil {
			connectMsg := &clients.Message{}
			connectMsg.MessageType = "DISCONNECT"
			connectMsg.SetSender(h)
			h.bus.Broadcast(connectMsg)
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

		var msg *clients.Message

		if h.version == "2.0" {
			msg = parse20(string(message))
		} else if h.version == "2.1" {
			msg = parse21(string(message))
		}

		if msg == nil {
			log.Println("Unable to process message: " + string(message))
			continue
		}

		if msg.MessageType == "VERSION" {
			h.version = msg.Message
			continue
		}

		if msg.MessageType == "PING" {
			h.SendMessage(&clients.Message{MessageType: "PONG"})
			continue
		}

		msg.SetSender(h)
		if msg.MessageType == "AUTH" {
			if msg.Message != os.Getenv("chatsharePSK") {
				log.Println("Authentication failed from: " + msg.Name)
				break
			}
			log.Println("Authentication success from: " + msg.Name)

			h.name = &msg.Name
			h.bus.AddClient(h)

			connectMsg := &clients.Message{}
			connectMsg.MessageType = "CONNECT"
			connectMsg.SetSender(h)
			err = h.conn.WriteMessage(1, []byte("WELCOME"))
			if err != nil {
				log.Println("Unable to send a message to: " + msg.Name + ". Closing connection")
				break
			}
			h.bus.Broadcast(connectMsg)
		} else {
			h.bus.Broadcast(msg)
		}
	}
}

func (h *handler) SendMessage(message *clients.Message) {
	if message.MessageType == "MESSAGE" {
		var msg []byte
		if h.version == "2.0" {
			msg = []byte(build20(message))
		} else {
			msg = build21(message)
		}
		err := h.conn.WriteMessage(1, msg)
		if err != nil {
			log.Println("Unable to send message: " + err.Error())
			return
		}
	}
}

func (h *handler) GetName() *string {
	return h.name
}
