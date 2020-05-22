package main

import (
	"github.com/madhouseminers/chatshare-server/internal/clients"
	"github.com/madhouseminers/chatshare-server/internal/discord"
	"github.com/madhouseminers/chatshare-server/internal/websocket"
	"log"
	"sync"
)

func main() {
	log.Println("Starting Chatshare service")

	ws := &sync.WaitGroup{}
	bus := clients.CreateBus()
	discord.CreateBot(bus)
	websocket.StartServer(bus, ws)

	ws.Wait()
}
