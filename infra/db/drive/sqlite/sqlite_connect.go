package sqlite

import (
	"errors"

	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/infra/db"
	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/infra/db/drive"
	"github.com/jinzhu/gorm"
)

func ConnectionSqlite() (*gorm.DB, error) {
	sqlite := NewSqlite()
	driver := drive.Drive(sqlite)
	if driver == "" {
		return nil, errors.New("driver is empty")
	}

	dns := sqlite.Connect()
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
