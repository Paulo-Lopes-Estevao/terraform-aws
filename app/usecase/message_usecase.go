package usecase

import (
	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/app/repository"
	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/entities"
)

type IMessageUsecase interface {
	SendMessageToUser(createMessageDto *CreateMessageDTO) (*MessageResponseDTO, error)
	GetUserBydId(from string) (*MessageResponseDTO, error)
}

type messageUsecase struct {
	MessageRepository repository.IMessageRepository
}

func NewMessageUseCase(messageRepo repository.IMessageRepository) IMessageUsecase {
	return &messageUsecase{
		MessageRepository: messageRepo,
	}
}

func (m *messageUsecase) SendMessageToUser(createMessageDto *CreateMessageDTO) (*MessageResponseDTO, error) {

	message, err := entities.NewMessage(createMessageDto.FromID, createMessageDto.To, createMessageDto.Content)
	if err != nil {
		return nil, err
	}

	sendMessage, err := m.MessageRepository.SendMessage(message)
	if err != nil {
		return nil, err
	}

	messageDTO := MessageResponseDTO{}
	return messageDTO.FromMessage(sendMessage), nil
}

func (m *messageUsecase) GetUserBydId(from string) (*MessageResponseDTO, error) {
	message, err := m.MessageRepository.FindByIdUser(from)
	if err != nil {
		return nil, err
	}

	messageDTO := MessageResponseDTO{}
	return messageDTO.FromMessage(message), nil
}
