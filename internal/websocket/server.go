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
}

type messageBus interface {
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
		log.Println("Got health check")
		writer.Write([]byte("OK"))
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
