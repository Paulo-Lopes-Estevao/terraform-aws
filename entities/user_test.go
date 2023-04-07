package entities_test

import (
	"testing"

	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/entities"
)

func TestNewUser(t *testing.T) {
	user, err := entities.NewUser("Lopes")

	if err != nil {
		t.Error(err)
	}

	if err := user.CheckIfUserExists("Lope"); err != nil {
		t.Error(err)
	}
}
