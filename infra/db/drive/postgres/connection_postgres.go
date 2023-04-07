package postgres

import (
	"errors"

	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/infra/db"
	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/infra/db/drive"
	"github.com/jinzhu/gorm"
)

func ConnectionPostgres() (*gorm.DB, error) {
	postgres := NewPostgres()
	driver := drive.Drive(postgres)
	if driver == "" {
		return nil, errors.New("driver is empty")
	}

	dns := postgres.Connect()
	if dns == "" {
		return nil, errors.New("dns is empty")
	}

	db, err := db.NewConnection(dns)
	if err != nil {
		return nil, err
	}
	dbConn := db.GetDatabaseConnection()
	if dbConn == nil {
		return nil, errors.New("dbConn is nil")
	}

	return dbConn, nil
}
