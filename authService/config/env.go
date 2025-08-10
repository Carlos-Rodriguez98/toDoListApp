package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	DBHOST     string
	DBPORT     int
	DBUser     string
	DBPassword string
	DBName     string

	ServerPort string
}

var AppConfig EnvConfig

func LoadEnv() {
	//Carga del archivo .env si existe
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar .env, usando variables de entorno")
	}

	AppConfig = EnvConfig{
		DBHOST:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		ServerPort: os.Getenv("PORT"),
	}
}
