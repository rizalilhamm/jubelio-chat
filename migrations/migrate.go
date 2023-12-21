package main

import (
	"fmt"
	"jubelio.com/chat/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"jubelio.com/chat/migrations/models"
)

func main() {
	db, err := createConnection()
	if err != nil {
		panic("Failed to connect database postgre")
	}

	log.Println("Migration: START")
	if err := db.AutoMigrate(modelTables...); err != nil {
		panic("Migration: " + err.Error())
	}
	log.Println("Migration: SUCCESS")
}

var modelTables []interface{} = []interface{} {
	&models.Users{},
	&models.Chats{},
	&models.Messages{},
}

func createConnection() (*gorm.DB, error) {
	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.GlobalEnv.PostgresHost,
		config.GlobalEnv.PostgresUser,
		config.GlobalEnv.PostgresPassword,
		config.GlobalEnv.PostgresDBName,
		config.GlobalEnv.PostgresPort,
		config.GlobalEnv.PostgresSSLMode,
	)
	return gorm.Open(postgres.Open(connection), &gorm.Config{})
}