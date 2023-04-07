package entities

import "errors"

type Message struct {
	ID      string
	From    string
	To      string
	Content string
	User    User
}

func NewMessage(from, to, content string) (*Message, error) {
	message := &Message{
		From:    from,
		To:      to,
		Content: content,
	}
	if err := message.isValid(); err != nil {
		return nil, err
	}
	return message, nil
}

func (m *Message) isValid() error {
	if m.Content == "" {
		return errors.New("o conteúdo é obrigatório")
	}
	if m.From == m.To {
		return errors.New("apenas permitido enviar menssagem para usuário diferente")
	}
	if m.To == "" {
		return errors.New("o destinatário é obrigatório")
	}
	if m.From == "" {
		return errors.New("o identificador é obrigatório")
	}
	return nil
}
