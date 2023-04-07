package usecase

import "github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/entities"

type CreateUserDTO struct {
	Name string
}

type LoginUserDTO struct {
	Name string
}

type UserResponse struct {
	ID   string
	Name string
}

type MessageResponseDTO struct {
	ID      string
	FromID    string `json:"from_id"`
	To      string
	Content string
}

type CreateMessageDTO struct {
	FromID    string	`json:"from_id"`
	To      string
	Content string
}

func (u *UserResponse) ToUser() *entities.User {
	return &entities.User{
		ID:   u.ID,
		Name: u.Name,
	}
}

func (m *MessageResponseDTO) ToMessage() *entities.Message {
	return &entities.Message{
		ID:      m.ID,
		FromID:    m.FromID,
		To:      m.To,
		Content: m.Content,
	}
}

func (u *UserResponse) FromUser(user *entities.User) *UserResponse {
	u.ID = user.ID
	u.Name = user.Name
	return u
}

func (m *MessageResponseDTO) FromMessage(message *entities.Message) *MessageResponseDTO {
	m.ID = message.ID
	m.FromID = message.FromID
	m.To = message.To
	m.Content = message.Content
	return m
}

func (u *UserResponse) FromUsers(users []*entities.User) []*UserResponse {
	var usersResponse []*UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, (&UserResponse{}).FromUser(user))
	}
	return usersResponse
}

func (m *MessageResponseDTO) FromMessages(messages []*entities.Message) []*MessageResponseDTO {
	var messagesResponse []*MessageResponseDTO
	for _, message := range messages {
		messagesResponse = append(messagesResponse, (&MessageResponseDTO{}).FromMessage(message))
	}
	return messagesResponse
}

