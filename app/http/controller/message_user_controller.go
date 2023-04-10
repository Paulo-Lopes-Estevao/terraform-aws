package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/app/usecase"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
)

type IMessageUserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	SendMessage(w http.ResponseWriter, r *http.Request)
	FindByMessageUser(w http.ResponseWriter, r *http.Request)
	ListAllUsersOntheNetwork(w http.ResponseWriter, r *http.Request)
}

type messageUserController struct {
	usecase.IMessageUsecase
	usecase.IUserUsecase
}

func NewMessageUserController(messageUsecase usecase.IMessageUsecase, userUsecase usecase.IUserUsecase) IMessageUserController {
	return &messageUserController{
		IMessageUsecase: messageUsecase,
		IUserUsecase:    userUsecase,
	}
}

func (m *messageUserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserDTO usecase.CreateUserDTO

	if err := json.NewDecoder(r.Body).Decode(&createUserDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := m.IUserUsecase.CreateUser(&createUserDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (m *messageUserController) Login(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := m.IUserUsecase.GetByNameUser(name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (m *messageUserController) SendMessage(w http.ResponseWriter, r *http.Request) {
	createMessageDTO := &usecase.CreateMessageDTO{}

	if err := json.NewDecoder(r.Body).Decode(createMessageDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	message, err := m.IMessageUsecase.SendMessageToUser(createMessageDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(message)
}

func (m *messageUserController) FindByMessageUser(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	if from == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	messages, err := m.IMessageUsecase.GetUserBydId(from)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(messages)
}

func (m *messageUserController) ListAllUsersOntheNetwork(w http.ResponseWriter, r *http.Request) {
	ctx := baggage.ContextWithoutBaggage(r.Context())

	trace := otel.Tracer("app/http/controller/message_user_controller.go")
	ctx, spanList := trace.Start(ctx, "ListAllUsersOntheNetwork")

	users, err := m.IUserUsecase.GetAllUser()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer spanList.End()


	ctx, spanDisplayResult := trace.Start(ctx, "DisplayResult")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
	defer spanDisplayResult.End()
}
