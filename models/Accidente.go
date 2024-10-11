package models

import (
	"time"

	"gorm.io/gorm"
)

type Accidente struct {
	gorm.Model
	Direccion         string `gorm:"not null"`
	Descripcion       string `gorm:"not null"`
	RequiereTranslado bool
	Fecha             time.Time
	Hora              time.Time
	PacienteID        uint
	AmbulanciaID      uint
	ReporteID         *uint     // Clave foránea que puede ser nula, indicando relación 0..1
	Reporte           *Reporte  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // Relación 0..1 con Reporte
	HospitalID        *uint     //puede ser null
	Hospital          *Hospital `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // Relación 1 a 1

}
