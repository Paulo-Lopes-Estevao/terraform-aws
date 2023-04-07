package service

import (
	"context"
	"fmt"

	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/app/grpc/pb"
	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/app/usecase"
	"go.opentelemetry.io/otel/trace"
	"github.com/uptrace/opentelemetry-go-extra/otelplay"
)

type MessageUser struct {
	pb.UnimplementedMessageUserServiceServer
	usecase.IMessageUsecase
	usecase.IUserUsecase
}

func NewMessageUserService(usecase usecase.IMessageUsecase, usecaseUser usecase.IUserUsecase) pb.MessageUserServiceServer {
	return &MessageUser{
		IMessageUsecase: usecase,
		IUserUsecase:    usecaseUser,
	}
}

func (m *MessageUser) CreateUser(ctx context.Context, request *pb.UserRequest) (*pb.UserResponse, error) {
	user, err := m.IUserUsecase.CreateUser(&usecase.CreateUserDTO{Name: request.GetName()})
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:   user.ID,
		Name: user.Name,
	}, nil
}

func (m *MessageUser) Login(ctx context.Context, request *pb.UserRequest) (*pb.UserResponse, error) {
	user, err := m.IUserUsecase.GetByNameUser(request.GetName())
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:   user.ID,
		Name: user.Name,
	}, nil
}

func (m *MessageUser) SendMessage(request *pb.MessageRequest, stream pb.MessageUserService_SendMessageServer) error {
	messageDTO := &usecase.CreateMessageDTO{}
	messageDTO.FromID = request.GetMessage().From
	messageDTO.To = request.GetMessage().To
	messageDTO.Content = request.GetMessage().Content

	message, err := m.IMessageUsecase.SendMessageToUser(messageDTO)
	if err != nil {
		return err
	}

	return stream.Send(&pb.MessageResponse{
		Id: message.ID,
		Message: &pb.Message{
			From:    message.FromID,
			To:      message.To,
			Content: message.Content,
		},
	})
}

func (m *MessageUser) FindByMessageUser(ctx context.Context, request *pb.GetUserById) (*pb.MessageResponse, error) {
	fmt.Println("trace", otelplay.TraceURL(trace.SpanFromContext(ctx)))
	messages, err := m.IMessageUsecase.GetUserBydId(request.GetFrom())
	if err != nil {
		return nil, err
	}

	return &pb.MessageResponse{
		Id: messages.ID,
		Message: &pb.Message{
			From:    messages.FromID,
			To:      messages.To,
			Content: messages.Content,
		},
	}, nil
}

func (m *MessageUser) ListAllUsersOntheNetwork(request *pb.ListUser, stream pb.MessageUserService_ListAllUsersOntheNetworkServer) error {
	users, err := m.IUserUsecase.GetAllUser()
	if err != nil {
		return err
	}

	for _, user := range users {
		err := stream.Send(&pb.UserResponse{
			Id:   user.ID,
			Name: user.Name,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
