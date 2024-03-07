package domains

import (
	"gorm.io/gorm"
)

func (EvaluacionDirector) TableName() string {
	return "evaluacion_director"
}

type EvaluacionDirector struct {
	gorm.Model
	// DirectorID                *uint  `gorm:"default:null"` //PersonalID del evaluador
	PersonalID                *uint  `gorm:"default:null"` //PersonalID del evaluador
	PeriodoID                 *uint  `gorm:"default:null"`
	ComPositivo               string `gorm:"default:null"`
	ObsSugerencia             string `gorm:"default:null"`
	DocenteID                 uint   `gorm:"not null"` //PersonalID del evaluado mejor en detalle
	DetalleEvaluacionDirector DetalleEvaluacionDirector
}
