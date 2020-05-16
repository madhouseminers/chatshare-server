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
	name    string
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

	b := &bot{session: discord, bus: bus, name: "Discord"}
	bus.AddClient(b)
	b.session.AddHandler(b.gotMessage)

	return b
}

func (b *bot) gotMessage(_ *discordgo.Session, message *discordgo.MessageCreate) {
	// Make sure we don't broadcast our own messages back
	if message.ChannelID == os.Getenv("discordChannel") && b.session.State.User.ID != message.Author.ID {
		b.bus.Broadcast(clients.CreateMessage("<"+message.Author.Username+"> "+message.Content, b))
	}
}

func (b *bot) SendMessage(message *clients.Message) {
	_, err := b.session.ChannelMessageSend(os.Getenv("discordChannel"), message.GetContent())
	if err != nil {
		log.Println("Unable to send discord message: " + err.Error())
	}
}

func (b *bot) GetName() *string {
	return &b.name
}
