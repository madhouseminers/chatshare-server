package clients

type Message struct {
	content string
	sender  Client
}

func CreateMessage(content string, sender Client) *Message {
	return &Message{content: "[" + *sender.GetName() + "] " + content, sender: sender}
}

func (m *Message) GetContent() string {
	return m.content
}

func (m *Message) GetSender() Client {
	return m.sender
}
