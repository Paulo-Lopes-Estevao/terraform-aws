package postgres

import (
	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/infra/db/drive"
	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/util"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type IPostgres interface {
	Connect() string
}

type Postgres struct {
	drive.IDrive
}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func (p *Postgres) Connect() string {
	env := util.LoadEnv()

	return "host=" + env.DBHOST + " port=" + env.DBPORT + " user=" + env.DBUSER + " dbname=" + env.DBNAME + " password=" + env.DBPASS + " sslmode=" + env.SSLMODE
}
