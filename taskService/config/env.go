package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
	JWTSecret  string
}

var AppConfig EnvConfig

func LoadEnv() {
	_ = godotenv.Load()

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil || port == 0 {
		port = 5432
	}

	AppConfig = EnvConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     port,
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		ServerPort: os.Getenv("PORT"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
	}

	if AppConfig.JWTSecret == "" {
		log.Println("ADVERTENCIA: JWT_SECRET no está definido")
	}
	if AppConfig.ServerPort == "" {
		AppConfig.ServerPort = "8080"
	}
}
