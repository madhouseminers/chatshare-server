package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/madhouseminers/chatshare-server/internal/clients"
	"log"
	"os"
)

type messageBus interface {
	AddClient(client clients.Client)
	RemoveClient(client clients.Client)
	Broadcast(message *clients.Message)
}

type bot struct {
	session *discordgo.Session
	name    *string
	bus     messageBus
}

func CreateBot(bus messageBus) *bot {
	discord, err := discordgo.New("Bot " + os.Getenv("discordToken"))
	if err != nil {
		log.Println("Error connecting to Discord: " + err.Error())
		return nil
	}
	err = discord.Open()
	if err != nil {
		log.Println("Error connecting to Discord: " + err.Error())
	}

	name := "Discord"
	b := &bot{session: discord, bus: bus, name: &name}
	bus.AddClient(b)
	b.session.AddHandler(b.gotMessage)
	_, err = b.session.ChannelMessageSend(os.Getenv("discordChannel"), "🤖 ChatShare 2.0 is connected")
	if err != nil {
		log.Println("Error sending connected message: " + err.Error())
	}

	return b
}

func (b *bot) gotMessage(_ *discordgo.Session, message *discordgo.MessageCreate) {
	// Make sure we don't broadcast our own messages back
	if message.ChannelID == os.Getenv("discordChannel") && b.session.State.User.ID != message.Author.ID {
		msg := &clients.Message{}
		msg.MessageType = "MESSAGE"
		msg.Name = message.Author.Username
		msg.Message = message.Content
		msg.SetSender(b)

		b.bus.Broadcast(msg)
	}
}

func (b *bot) SendMessage(message *clients.Message) {
	textToSend := ""
	if message.MessageType == "CONNECT" {
		textToSend = "👋 " + *message.GetSender().GetName() + " has connected"
	} else if message.MessageType == "DISCONNECT" {
		textToSend = "👋 " + *message.GetSender().GetName() + " has disconnected"
	} else if message.MessageType == "JOIN" {
		textToSend = "➡ " + message.Name + " has joined " + *message.GetSender().GetName()
	} else if message.MessageType == "LEAVE" {
		textToSend = "⬅ " + message.Name + " has left " + *message.GetSender().GetName()
	} else if message.MessageType == "MESSAGE" {
		textToSend = "[" + *message.GetSender().GetName() + "] <" + message.Name + "> " + message.Message
	}

	if len(textToSend) == 0 {
		return
	}

	_, err := b.session.ChannelMessageSend(os.Getenv("discordChannel"), textToSend)
	if err != nil {
		log.Println("Unable to send discord message: " + err.Error())
	}
}

func (b *bot) GetName() *string {
	return b.name
}
