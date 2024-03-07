package domains

import (
	"gorm.io/gorm"
)

func (Evaluacion) TableName() string {
	return "evaluacion"
}

type Evaluacion struct {
	gorm.Model
	AlumnoID          uint   `gorm:"not null"`
	PersonalID        uint   `gorm:"not null"` //evaluado
	PeriodoID         uint   `gorm:"not null"`
	Parcial           uint8  `gorm:"not null"`
	GrupoMateriaID    uint   `gorm:"not null"`
	ComPositivo       string `gorm:"default:null"`
	ObsSugerencia     string `gorm:"default:null"`
	MateriaID         uint   `gorm:"not null"`
	DetalleEvaluacion DetalleEvaluacion
}
