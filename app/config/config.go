package config

import (
	"flag"
	"github.com/gofiber/fiber/v3/log"
	"github.com/joho/godotenv"
	"os"
)

type Database struct {
	Host     string
	User     string
	Port     string
	Password string
	Database string
}
type Config struct {
	DB Database
}

func GetConfig() *Config {

	fileFlag := flag.String("env", "", "File .env path absolute")

	var err error
	if *fileFlag != "" {
		err = godotenv.Load(*fileFlag)
	} else {
		err = godotenv.Load()
	}
	if err != nil {
		log.Fatal("Error pada saat load .env " + err.Error())
	}
	return &Config{
		DB: Database{
			User:     os.Getenv("DATABASE_USER"),
			Database: os.Getenv("DATABASE_NAME"),
			Password: os.Getenv("DATABASE_PASS"),
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     os.Getenv("DATABASE_PORT"),
		},
	}
}
