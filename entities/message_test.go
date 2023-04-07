package entities_test

import (
	"testing"

	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/entities"
)

func TestNewMessage(t *testing.T) {
	message, err := entities.NewMessage("1","2","Hello Word")
	if err != nil {
		t.Error(err)
	}

	t.Log(message)
}
