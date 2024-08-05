package database

import (
	"gateway/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func ConnectToDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DBConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = DBConn.Migrator().AutoMigrate(&models.Port{})
	if err != nil {

		return nil, err
	}

	return DBConn, nil
}

