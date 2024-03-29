package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/madhouseminers/chatshare-server/internal/clients"
	"log"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type messageBus interface {
	Direct(message *clients.Message, clientName string)
	AddClient(client clients.Client)
	RemoveClient(client clients.Client)
	Broadcast(message *clients.Message)
}

type httpServer struct {
	bus messageBus
}

func StartServer(bus messageBus, ws *sync.WaitGroup) *httpServer {
	server := &http.Server{
		Addr: ":8080",
	}
	h := &httpServer{bus: bus}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("OK"))
		if err != nil {
			log.Println("Error responding to health check: " + err.Error())
		}
	})
	http.HandleFunc("/ws", h.upgradeHandler)
	ws.Add(1)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Println("Unable to close server: " + err.Error())
		}
		ws.Done()
	}()
	return h
}

func (h *httpServer) upgradeHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading request: " + err.Error())
		return
	}
	createHandler(conn, h.bus)
}
