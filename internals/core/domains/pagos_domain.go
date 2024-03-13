package domains

import (
	"time"

	"gorm.io/gorm"
)

type Pago struct {
	gorm.Model
	EstudianteID uint      `gorm:"not null"`
	PeriodoID    uint      `gorm:"not null"`
	TipoPagoID   uint      `gorm:"not null"`
	Fecha        time.Time `gorm:"not null"`
}
