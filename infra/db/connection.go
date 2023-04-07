package db

import (
	"os"
	"sync"

	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/entities"
	"github.com/jinzhu/gorm"
)

type Connection interface {
	GetDatabaseConnection() *gorm.DB
}

type connection struct {
	db *gorm.DB
}

var (
	db   *gorm.DB
	once = &sync.Once{}
)

func (c *connection) GetDatabaseConnection() *gorm.DB {
	return c.db
}

func NewConnection(dns string) (Connection, error) {

	once.Do(func() {
		var err error
		db, err = gorm.Open(os.Getenv("DB_CONNECTION"), dns)
		if err != nil {
			panic(err)
		}

		db.AutoMigrate(&entities.Message{}, &entities.User{})
	})

	return &connection{db: db}, nil
}
