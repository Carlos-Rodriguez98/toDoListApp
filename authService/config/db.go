package config

import (
	"fmt"
	"log"
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

	var err error                                           //Declaro variable para captura de errores
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{}) //Abre la conexión, env.dbURL contiene la información.
	if err != nil {                                         //Generar un mensaje de error si err es diferente de nUll
		log.Fatal("Error conectando a la base de datos:", err)
	}
	log.Println("Conexión a base de datos exitosa")

	//Automigración
	if err := DB.AutoMigrate(&models.Usuario{}); err != nil {
		log.Fatal("Error en la migración: ", err)
	}
	log.Println("Migración completada")
}
