package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"strconv"

)

type Env struct {
	SupabaseUrl string
	SupabaseKey string
	
	PostgresHost         string
	PostgresUser         string
	PostgresPassword     string
	PostgresDBName       string
	PostgresPort         uint16
	PostgresSSLMode      string
	PostgresMaxIdleConns int
	PostgresMaxOpenConns int
	PostgresMaxLifeTime  int

	BasicAuthUsername string
	BasicAuthPassword string
}

var GlobalEnv Env

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	var ok bool

	GlobalEnv.SupabaseUrl = os.Getenv("SUPABASE_URL")
	GlobalEnv.SupabaseKey = os.Getenv("SUPABASE_KEY")

	GlobalEnv.PostgresHost, ok = os.LookupEnv("POSTGRES_DB_HOST")
	if !ok {
		panic("missing POSTGRES_DB_HOST environment")
	}

	if postgresPort, err := strconv.Atoi(os.Getenv("POSTGRES_DB_PORT")); err != nil {
		panic("missing POSTGRES_DB_PORT environment")
	} else {
		GlobalEnv.PostgresPort = uint16(postgresPort)
	}

	GlobalEnv.PostgresUser, ok = os.LookupEnv("POSTGRES_DB_USER")
	if !ok {
		panic("missing POSTGRES_DB_USER environment")
	}

	GlobalEnv.PostgresPassword, ok = os.LookupEnv("POSTGRES_DB_PASSWORD")
	if !ok {
		panic("missing POSTGRES_DB_PASSWORD environment")
	}

	GlobalEnv.PostgresDBName, ok = os.LookupEnv("POSTGRES_DB_NAME")
	if !ok {
		panic("missing POSTGRES_DB_NAME environment")
	}

	GlobalEnv.PostgresSSLMode, ok = os.LookupEnv("POSTGRES_DB_SSLMODE")
	if !ok {
		panic("missing POSTGRES_DB_SSLMODE environment")
	}
	GlobalEnv.PostgresDBName, ok = os.LookupEnv("BASIC_AUTH_USERNAME")
	if !ok {
		panic("missing BASIC_AUTH_USERNAME environment")
	}

	GlobalEnv.PostgresSSLMode, ok = os.LookupEnv("BASIC_AUTH_PASSWORD")
	if !ok {
		panic("missing BASIC_AUTH_PASSWORD environment")
	}

}