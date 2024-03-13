package domains

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	Pregunta     string `gorm:"not null"`
	ClasifPregID uint   `gorm:"not null"`
	Tipo         int    `gorm:"not null"`
}
