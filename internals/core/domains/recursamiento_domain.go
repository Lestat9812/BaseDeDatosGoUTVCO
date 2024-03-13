package domains

import "gorm.io/gorm"

type Recursamiento struct {
	gorm.Model
	CalificacionID uint `gorm:"not null"`
	PeriodoID      uint `gorm:"not null"`
	PersonalID     uint `gorm:"not null"`
	Calificacion   int  `gorm:"not null"` //index en base original ?????
}
