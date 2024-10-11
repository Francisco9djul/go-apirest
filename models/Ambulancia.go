package models

import (
	"gorm.io/gorm"
)

type Ambulancia struct {
	gorm.Model
	Patente      string `gorm:"not null;unique"`
	Inventario   bool
	Vtv          bool
	Seguro       bool
	Estado       TipoEstado  `gorm:"type:varchar(5);not null; default:'alta'"` // Relacionado con la tabla TipoEstado
	ChoferID     uint        `gorm:"unique"`
	ParamedicoID uint        `gorm:"unique"`
	Accidentes   []Accidente `gorm:"foreignKey:AmbulanciaID;"` // Relación con accidentes
}
