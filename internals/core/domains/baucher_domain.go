package domains

import "gorm.io/gorm"

type Baucher struct {
	gorm.Model
	AlumnoID      uint    `gorm:"not null"`
	EstudianteID  uint    `gorm:"not null"`
	PeriodoID     uint    `gorm:"not null"`
	CarreraID     uint    `gorm:"not null"`
	Fecha         string  `gorm:"not null"`
	Extension     string  `gorm:"not null"`
	NumReferencia string  `gorm:"not null"`
	Observaciones *string `gorm:"type:varchar(250);default:null"`
	EDID          uint    `gorm:"not null"`
}
