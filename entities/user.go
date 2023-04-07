package entities

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID   string
	Name string
}

func NewUser(name string) (*User, error) {
	user := &User{
		Name: name,
	}
	user.ID = uuid.New().String()
	if err := user.isValid(); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) isValid() error {
	if u.Name == "" {
		return errors.New("nome é obrigatório")
	}
	return nil
}

func (u *User) CheckIfUserExists(name string) error {
	if u.Name == name {
		return errors.New("nome do usuário já existe")
	}
	return nil
}
