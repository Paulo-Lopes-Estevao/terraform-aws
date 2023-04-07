package repository

import (
	"errors"

	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/entities"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (u *UserRepository) Save(user *entities.User) error {
	return u.DB.Save(user).Error
}

func (u *UserRepository) FindByName(name string) (*entities.User, error) {
	user := &entities.User{}
	u.DB.Where("name = ?", name).First(user)
	return user, errors.New("usuario nao encontrado")
}

func (u *UserRepository) FindAll() ([]*entities.User, error) {
	users := []*entities.User{}
	err := u.DB.Find(&users).Error
	return users, err
}
