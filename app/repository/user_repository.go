package repository

import "github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/entities"


type IUserRepository interface{
	Save(user *entities.User) error
	FindByName(name string) (*entities.User, error)
	FindAll() ([]*entities.User, error)
}