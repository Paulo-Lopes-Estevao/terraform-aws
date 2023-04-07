package main

import (
	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/app/grpc/server"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	server.ServerGrpc()
}
