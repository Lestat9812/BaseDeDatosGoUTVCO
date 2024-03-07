package domains

import "gorm.io/gorm"

func (Autoevaluacion) TableName() string {
	return "autoevaluaciones"
}

type Autoevaluacion struct {
	gorm.Model
	PersonalID            uint    `gorm:"not null"`
	PeriodoID             uint    `gorm:"not null"`
	ComPositivo           *string `gorm:"type:text;default:null"`
	ObsSugerencia         *string `gorm:"type:text;default:null"`
	MateriaID             uint    `gorm:"not null"`
	DetalleAutoevaluacion DetalleAutoevaluacion
}
