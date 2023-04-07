package server

import (
	"fmt"
	"log"
	"net"

	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/app/grpc/dependency"
	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/app/grpc/pb"
	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/infra/db/drive/sqlite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func ServerGrpc() {
	fmt.Println("server grpc")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	db, err := sqlite.ConnectionSqlite()

	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	serviceDependency := dependency.Dependency(db)

	grpcServer := grpc.NewServer()
	pb.RegisterMessageUserServiceServer(grpcServer, serviceDependency)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalln(err)
	}

}
