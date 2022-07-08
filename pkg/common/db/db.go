package db

import (
	"github.com/erica7dev/fiberapi/pkg/common/models"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"log"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}