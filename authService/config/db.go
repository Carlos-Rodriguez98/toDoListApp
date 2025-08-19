package config

import (
	"fmt"
	"log"
	"time"
	"toDoListApp/authService/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB //Declaración de variable global (puntero a gorm.DB)

// Función para establecer conexión a BD
func ConnectDatabase() {
	//Construcción de URL de BD
	DSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Bogota",
		AppConfig.DBHost, AppConfig.DBUser, AppConfig.DBPassword, AppConfig.DBName, AppConfig.DBPort,
	)

	var err error //Declaro variable para captura de errores
	//Intento de conexión a la base de datos (hasta 5 intentos cada 2 seg)
	for i := 1; i <= 5; i++ {
		DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
		if err == nil {
			log.Print("Conexión exitosa a la base de datos")

			//Ejecución de automigración
			if err := DB.AutoMigrate(&models.Usuario{}); err != nil {
				log.Fatal("Error en la migración: ", err)
			}
			log.Println("Migración completada")
			return
		}
		log.Printf("Intento %d: error conectando a la base de datos: %v", i, err)
		time.Sleep(2 * time.Second)
	}
	log.Fatal("No se pudo conectar a la base de datos después de varios intentos")
}
