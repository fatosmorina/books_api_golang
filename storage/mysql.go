package storage

import (
	"books_api/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var err error

	DB, err = gorm.Open(mysql.Open(config.GetDBUrl()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database. error: %v", err)
	}
	log.Println("Successfully connected database")

	return DB
}
