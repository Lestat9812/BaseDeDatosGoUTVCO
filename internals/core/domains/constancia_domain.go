package domains

import (
	"time"

	"gorm.io/gorm"
)

type Constancia struct {
	gorm.Model
	AlumnoID          uint      `gorm:"not null"`
	PeriodoID         uint      `gorm:"not null"`
	Grupo             string    `gorm:"not null"`
	ClaveConstancia   string    `gorm:"not null"`
	Telefono          string    `gorm:"not null"`
	Correo            string    `gorm:"not null"`
	Fecha             time.Time `gorm:"not null"`
	Status            int       `gorm:"not null"`
	Observaciones     string    `gorm:"not null"`
	NumReferencia     string    `gorm:"not null"`
	EstadoPago        int       `gorm:"not null"`
	FechaS            *time.Time
	ObservacionesPago string `gorm:"not null"`
	ConstanciaEnv     int    `gorm:"not null"`
}
