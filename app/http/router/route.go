package router

import (
	"net/http"

	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/app/http/controller"
)

type Router struct {
	controller.IMessageUserController
	SeverMux *http.ServeMux
}

func NewRouter(messageUserController controller.IMessageUserController, serverMux *http.ServeMux) *Router {
	return &Router{
		IMessageUserController: messageUserController,
		SeverMux:               serverMux,
	}
}

func (r *Router) Router() *http.ServeMux {

	r.SeverMux.HandleFunc("/user", r.IMessageUserController.CreateUser)
	r.SeverMux.HandleFunc("/login", r.IMessageUserController.Login)
	r.SeverMux.HandleFunc("/message", r.IMessageUserController.SendMessage)
	r.SeverMux.HandleFunc("/message/user", r.IMessageUserController.FindByMessageUser)
	r.SeverMux.HandleFunc("/message/user/all", r.IMessageUserController.ListAllUsersOntheNetwork)

	return r.SeverMux
}
