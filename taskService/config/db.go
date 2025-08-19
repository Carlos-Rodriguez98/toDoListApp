package config

import (
	"fmt"
	"log"
	"time"
	"toDoListApp/taskService/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	DSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Bogota",
		AppConfig.DBHost, AppConfig.DBUser, AppConfig.DBPassword, AppConfig.DBName, AppConfig.DBPort,
	)

	var err error
	for i := 1; i <= 5; i++ {
		DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
		if err == nil {
			log.Println("Conexión exitosa a la base de datos (taskService)")
			if err := DB.AutoMigrate(&models.Tarea{}); err != nil {
				log.Fatal("Error en la migración de Tarea:", err)
			}
			return
		}
		log.Printf("Intento %d: error conectando a la base de datos: %v", i, err)
		time.Sleep(2 * time.Second)
	}
	log.Fatal("No se pudo conectar a la base de datos después de varios intentos")
}
