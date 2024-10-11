package models

import (
	"gorm.io/gorm"
)

type Paramedico struct {
	gorm.Model
	DNI            string       `gorm:"not null;unique"`
	NombreCompleto string       `gorm:"not null"`
	Estado         TipoEstado   `gorm:"type:varchar(5);not null; default:'alta'"` // Relacionado con la tabla TipoEstado
	Ambulancias    []Ambulancia `gorm:"foreignKey:ParamedicoID;"`                 // Relación con ambulancias
}
