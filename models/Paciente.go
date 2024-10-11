package models

import (
	"gorm.io/gorm"
)

type Paciente struct {
	gorm.Model
	DNI            string      `gorm:"not null;unique"`
	NombreCompleto string      `gorm:"not null"`
	Telefono       string      `gorm:"type:varchar(15)"`
	Accidentes     []Accidente `gorm:"foreignKey:PacienteID;"` // Relación con accidentes

}
