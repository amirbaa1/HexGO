package server

import (
	"book/internal/core/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	database, err := gorm.Open(
		postgres.Open("host=localhost port=5432 user=postgres dbname=bookGO sslmode=disable password=amir$$1379"),
		&gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = database.AutoMigrate(&model.Book{}, &model.Author{})
	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
	db = database
}

func GetDB() *gorm.DB {
	return db
}
