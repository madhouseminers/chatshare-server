package websocket

import (
	"github.com/madhouseminers/chatshare-server/internal/clients"
	"strings"
)

func parse20(message string) *clients.Message {

	msg := &clients.Message{}

	if message[0] == '<' {
		// <playerX> xxxx
		msg.MessageType = "MESSAGE"

		messagePieces := strings.SplitN(message, "> ", 2)

		msg.Name = strings.Replace(messagePieces[0], "<", "", 1)
		msg.Message = messagePieces[1]
	} else if (strings.Index(message, "has left")) != -1 {
		// playerX has left
		msg.MessageType = "LEAVE"
		msg.Name = strings.Replace(message, " has left", "", 1)
	} else if (strings.Index(message, "has joined")) != -1 {
		// playerX has joined
		msg.MessageType = "JOIN"
		msg.Name = strings.Replace(message, " has joined", "", 1)
	} else if (strings.Index(message, "VERSION")) == 0 {
		messagePieces := strings.SplitN(message, "::", 2)

		msg.MessageType = "VERSION"
		msg.Message = messagePieces[1]
	} else {
		// server name::auth token
		msg.MessageType = "AUTH"

		messagePieces := strings.SplitN(message, "::", 2)

		msg.Name = messagePieces[0]
		msg.Message = messagePieces[1]
	}

	return msg
}

func build20(msg *clients.Message) string {
	return "[" + *msg.GetSender().GetName() + "] <" + msg.Name + "> " + msg.Message
}
