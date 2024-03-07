package domains

import "gorm.io/gorm"

func (DetalleAutoevaluacion) TableName() string {
	return "detalle_autoevaluacion"
}

type DetalleAutoevaluacion struct {
	gorm.Model
	AutoevaluacionID *uint `gorm:"default:null"`
	QuestionID       *uint `gorm:"default:null"`
	Calificacion     *uint `gorm:"default:null"`
	PersonalID       uint  `gorm:"not null"` //redundante?
	MateriaID        uint  `gorm:"not null"` //redundante?
	PeriodoID        uint  `gorm:"not null"` //redundante?
}
