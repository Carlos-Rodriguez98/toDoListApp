package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB //Declaración de variable global (puntero a gorm.DB)

// Función para establecer conexión a BD
func ConnectDatabase() {
	var err error                                                 //Declaro variable para captura de errores
	DB, err = gorm.Open(postgres.Open(env.dbURL), &gorm.Config{}) //Abre la conexión, env.dbURL contiene la información.
	if err != nil {                                               //Generar un mensaje de error si err es diferente de nUll
		log.Fatal("Error conectando a la base de datos:", err)
	}
	log.Println("Conexión a base de datos exitosa")
}
