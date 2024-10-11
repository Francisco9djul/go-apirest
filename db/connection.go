package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Variable global para la conexión
var DB *gorm.DB

// connection string
var DNS = "host=localhost user=admin password=admin2001 dbname=go-apirest port=5432"

func DBConnection() {

	var err error
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v\n", err)
	}

	fmt.Println("Conexión exitosa a la base de datos")
}
