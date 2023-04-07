package entities

import (
	"errors"

	"github.com/google/uuid"
)

type Message struct {
	ID      string
	FromID  string `json:"from_id"`
	To      string
	Content string
}

func NewMessage(from, to, content string) (*Message, error) {
	message := &Message{
		FromID:    from,
		To:      to,
		Content: content,
	}
	message.ID = uuid.New().String()
	if err := message.isValid(); err != nil {
		return nil, err
	}
	return message, nil
}

func (m *Message) isValid() error {
	if m.Content == "" {
		return errors.New("o conteúdo é obrigatório")
	}
	if m.FromID == m.To {
		return errors.New("apenas permitido enviar menssagem para usuário diferente")
	}
	if m.To == "" {
		return errors.New("o destinatário é obrigatório")
	}
	if m.FromID == "" {
		return errors.New("o identificador é obrigatório")
	}
	return nil
}
