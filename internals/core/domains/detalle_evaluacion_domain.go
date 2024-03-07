package domains

import "gorm.io/gorm"

func (DetalleEvaluacion) TableName() string {
	return "detalle_evaluacion"
}

type DetalleEvaluacion struct {
	gorm.Model
	EvaluacionID uint   `gorm:"not null"`
	QuestionID   uint   `gorm:"not null"`
	Calificacion string `gorm:"not null"`
	PersonalID   uint   `gorm:"not null"`
	MateriaID    uint   `gorm:"not null"`
	AlumnoID     uint   `gorm:"not null"`
	GrupoID      uint   `gorm:"not null"`
}
