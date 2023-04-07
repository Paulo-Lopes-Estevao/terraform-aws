package repository

import "github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/entities"

type IMessageRepository interface {
	SendMessage(message *entities.Message) (*entities.Message, error)
	FindByIdUser(from string) (*entities.Message, error)
}
