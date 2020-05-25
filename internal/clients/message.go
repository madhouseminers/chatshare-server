package clients

type Message struct {
	direct bool
	sender Client

	Name        string `json:"name"`
	Message     string `json:"message"`
	MessageType string `json:"type"`
	Sender      string `json:"sender"`
}

func (m *Message) GetContent() string {
	if m.direct {
		return m.Message
	} else {
		return "[" + *m.sender.GetName() + "] " + m.Message
	}
}

func (m *Message) SetDirect() *Message {
	m.direct = true
	return m
}

func (m *Message) GetSender() Client {
	return m.sender
}

func (m *Message) SetSender(sender Client) {
	m.sender = sender
	if m.sender.GetName() != nil {
		m.Sender = *sender.GetName()
	}
}
