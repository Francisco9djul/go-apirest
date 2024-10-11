package models

import (
	"gorm.io/gorm"
)

type Hospital struct {
	gorm.Model

	Nombre     string `gorm:"not null"`
	Direccion  string
	Accidentes []Accidente `gorm:"foreignKey:PacienteID;"` // Relación con accidentes
}
