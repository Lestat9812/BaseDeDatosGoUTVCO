package domains

import "gorm.io/gorm"

func (DetalleEvaluacionDirector) TableName() string {
	return "detalle_evaluacion_director"
}

type DetalleEvaluacionDirector struct {
	gorm.Model
	EvaluacionDirectorID *uint `gorm:"default:null"`
	QuestionID           *uint `gorm:"default:null"`
	Calificacion         *uint `gorm:"default:null"`
	PersonalID           *uint `gorm:"default:null"` //evaluado
	PeriodoID            uint  `gorm:"not null"`
}
