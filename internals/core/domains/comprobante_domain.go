package domains

import "gorm.io/gorm"

type Comprobante struct {
	gorm.Model
	AlumnoID  uint   `gorm:"not null"`
	GrupoID   uint   `gorm:"not null"`
	Parcial   int    `gorm:"not null"`
	Folio     string `gorm:"not null"`
	PeriodoID uint   `gorm:"not null"`
}
