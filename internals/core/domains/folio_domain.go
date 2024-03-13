package domains

import (
	"time"

	"gorm.io/gorm"
)

type Folio struct {
	gorm.Model
	TsuIng           string `gorm:"not null"`
	Folio            int    `gorm:"not null"`
	Libro            int    `gorm:"not null"`
	Foja             int    `gorm:"not null"`
	AlumnoID         uint   `gorm:"not null"`
	CarreraID        uint   `gorm:"not null"`
	FolioDocumentoID uint   `gorm:"not null"`
	Posicion         int    `gorm:"not null"`
	// TituloID         uint      `gorm:"not null"` //??
	Promedio float64   `gorm:"not null"`
	Fecha    time.Time `gorm:"not null"`
}
