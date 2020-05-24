package websocket

import (
	"encoding/json"
	"github.com/madhouseminers/chatshare-server/internal/clients"
	"log"
)

func parse21(message string) *clients.Message {

	msg := &clients.Message{}
	err := json.Unmarshal([]byte(message), &msg)
	if err != nil {
		return nil
	}

	return msg
}

func build21(msg *clients.Message) []byte {
	output, err := json.Marshal(msg)
	if err != nil {
		log.Println("Error packaging message: " + err.Error())
		return nil
	}

	return output
}
