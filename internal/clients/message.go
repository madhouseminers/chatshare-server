package clients

type Message struct {
	direct  bool
	content string
	sender  Client
}

func CreateMessage(content string, sender Client) *Message {
	return &Message{direct: false, content: content, sender: sender}
}

func (m *Message) GetContent() string {
	if m.direct {
		return m.content
	} else {
		return "[" + *m.sender.GetName() + "] " + m.content
	}
}

func (m *Message) SetDirect() *Message {
	m.direct = true
	return m
}

func (m *Message) GetSender() Client {
	return m.sender
}
