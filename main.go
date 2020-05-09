package main

import (
	"github.com/madhouseminers/chatshare-server/internal/clients"
	"github.com/madhouseminers/chatshare-server/internal/discord"
	"github.com/madhouseminers/chatshare-server/internal/websocket"
	"sync"
)

func main() {
	ws := &sync.WaitGroup{}
	bus := clients.CreateBus()
	websocket.StartServer(bus, ws)
	discord.CreateBot(bus)

	ws.Wait()
}
