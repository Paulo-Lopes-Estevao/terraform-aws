package dependency

import (
	"net/http"

	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/app/http/controller"
	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/app/http/router"
	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/app/usecase"
	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/infra/repository"
	"github.com/jinzhu/gorm"
)

func Dependency(db *gorm.DB, serverMux *http.ServeMux) *http.ServeMux {
	UserRepository := repository.NewUserRepository(db)
	UserUsecase := usecase.NewUserUseCase(UserRepository)

	MessageRepository := repository.NewMessageRepository(db)
	MessageUsecase := usecase.NewMessageUseCase(MessageRepository)

	messageUserController := controller.NewMessageUserController(MessageUsecase, UserUsecase)

	r := router.NewRouter(messageUserController, serverMux)

	return r.Router()
}
