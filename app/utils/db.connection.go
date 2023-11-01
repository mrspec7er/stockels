package utils

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	dbConfig := os.Getenv("DB_CONFIG")
	dbConnection, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	
	if err != nil {
		panic(err)
	}

	return dbConnection
}