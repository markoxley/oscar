package common

type Message struct {
	source string
	text   string
}

func NewMessage(s, t string) Message {
	return Message{s, t}
}

func (m Message) Source() string {
	return m.source
}

func (m Message) Text() string {
	return m.text
}
