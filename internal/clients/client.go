package clients

type Client interface {
	GetName() *string
	SendMessage(message *Message)
}
