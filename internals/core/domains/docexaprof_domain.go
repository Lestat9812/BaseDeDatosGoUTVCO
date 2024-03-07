package domains

import (
	"time"

	"gorm.io/gorm"
)

func (DocExaProf) TableName() string {
	return "docexaprof"
}

type DocExaProf struct {
	gorm.Model
	AlumnoID uint      `gorm:"not null"`
	Docest   string    `gorm:"not null"` //????
	Periodo  uint      `gorm:"not null"`
	Fecha    time.Time `gorm:"not null"`
	Ext      string    `gorm:"not null"`
}
