package db

import (
	"log"

	"gorm.io/gorm"
)

func ConnectDatabase(dial gorm.Dialector) *gorm.DB {
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
