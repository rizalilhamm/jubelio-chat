package databases

import (
	"fmt"
	"time"

	"jubelio.com/chat/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgre() *gorm.DB {

	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.GlobalEnv.PostgresHost,
		config.GlobalEnv.PostgresUser,
		config.GlobalEnv.PostgresPassword,
		config.GlobalEnv.PostgresDBName,
		config.GlobalEnv.PostgresPort,
		config.GlobalEnv.PostgresSSLMode,
	)

	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database postgre")
	}

	postgresDb, err := db.DB()
	if err != nil {
		panic("Failed to create pool connection database postgre")
	}

	PostgresMaxLifeTime := time.Duration(config.GlobalEnv.PostgresMaxLifeTime)
	postgresDb.SetMaxOpenConns(config.GlobalEnv.PostgresMaxOpenConns)
	postgresDb.SetMaxIdleConns(config.GlobalEnv.PostgresMaxIdleConns)
	postgresDb.SetConnMaxLifetime(PostgresMaxLifeTime * time.Second)
	return db
}
