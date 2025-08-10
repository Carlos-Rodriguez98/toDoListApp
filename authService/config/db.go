package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB //Declaración de variable global (puntero a gorm.DB)

// Función para establecer conexión a BD
func ConnectDatabase(cfg EnvConfig) {
	//Construcción de URL de BD
	DSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Bogota",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	var err error                                           //Declaro variable para captura de errores
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{}) //Abre la conexión, env.dbURL contiene la información.
	if err != nil {                                         //Generar un mensaje de error si err es diferente de nUll
		log.Fatal("Error conectando a la base de datos:", err)
	}
	log.Println("Conexión a base de datos exitosa")
}
