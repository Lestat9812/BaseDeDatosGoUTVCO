package domains

import (
	"gorm.io/gorm"
)

func (Nota) TableName() string {
	return "notas"
}

type Nota struct {
	gorm.Model
	AlumnoID  uint   `gorm:"not null"`
	PeriodoID uint   `gorm:"not null"`
	Matricula string `gorm:"not null"`
	Nota      string `gorm:"not null"`
}
