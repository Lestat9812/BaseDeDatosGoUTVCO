package domains

import (
	"time"

	"gorm.io/gorm"
)

func (DocEstudiante) TableName() string {
	return "docestudiante"
}

type DocEstudiante struct {
	gorm.Model
	EstudianteID uint      `gorm:"not null"`
	Docest       string    `gorm:"not null"` //???
	Ext          string    `gorm:"not null"`
	Fecha        time.Time `gorm:"not null"`
	Periodo      uint      `gorm:"not null"`
}
