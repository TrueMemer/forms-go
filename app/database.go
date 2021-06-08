package app

import (
	"forms/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Database struct {
	database *gorm.DB
}

func (database *Database) InitDatabase() {
	db, _ := gorm.Open(postgres.Open(os.Getenv("DATABASE_DSN")), &gorm.Config{SkipDefaultTransaction: true})

	db.AutoMigrate(models.User{}, models.Group{}, models.Form{})

	database.database = db
}

func (database *Database) GetDatabase() *gorm.DB {
	return database.database
}