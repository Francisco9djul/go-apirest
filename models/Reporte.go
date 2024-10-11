package models

import (
	"time"

	"gorm.io/gorm"
)

type Reporte struct {
	gorm.Model

	Descripcion string `gorm:"not null"`
	FechayHora  time.Time
}
