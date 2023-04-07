package repository

import (
	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/app/repository"
	"github.com/jinzhu/gorm"
)

type IUserRepositoryContract interface {
	UserContract() repository.IUserRepository
}

type IMessageRepositoryContract interface {
	MessageContract() repository.IMessageRepository
}

type IContract interface {
	IUserRepositoryContract
	IMessageRepositoryContract
}

type Contract struct {
	DB *gorm.DB
}

func NewContract(db *gorm.DB) *Contract {
	return &Contract{DB: db}
}

func (c *Contract) UserContract() repository.IUserRepository {
	return NewUserRepository(c.DB)
}

func (c *Contract) MessageContract() repository.IMessageRepository {
	return NewMessageRepository(c.DB)
}