package sqlite

import (
	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/infra/db/drive"
	_ "github.com/mattn/go-sqlite3"
)


type ISqlite interface {
	Connect() string
}

type Sqlite struct {
	drive.IDrive
}

func NewSqlite() *Sqlite {
	return &Sqlite{}
}

func (s *Sqlite) Connect() string {
	return "sqlite3.db"
}