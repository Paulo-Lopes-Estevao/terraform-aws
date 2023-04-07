package repository

import (
	"log"

	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/entities"
	"github.com/jinzhu/gorm"
)

type MessageRepository struct {
	Db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{Db: db}
}

func (m *MessageRepository) SendMessage(message *entities.Message) (*entities.Message, error) {
	err := m.Db.Save(message).Error
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (m *MessageRepository) FindByIdUser(from string) (*entities.Message, error) {
	message := &entities.Message{}
	log.Println("FROM: ", from)
	err := m.Db.Find(message, "from_id = ?", from).Error
	log.Println("ERROR: ", err)
	return message, err
}
