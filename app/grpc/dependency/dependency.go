package dependency

import (
	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/app/grpc/service"
	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/app/usecase"
	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/infra/repository"
	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/app/grpc/pb"
	"github.com/jinzhu/gorm"
)

func Dependency(db *gorm.DB) pb.MessageUserServiceServer {
	UserRepository := repository.NewUserRepository(db)
	UserUsecase := usecase.NewUserUseCase(UserRepository)

	MessageRepository := repository.NewMessageRepository(db)
	MessageUsecase := usecase.NewMessageUseCase(MessageRepository)

	Service := service.NewMessageUserService(MessageUsecase, UserUsecase)

	return Service
}
