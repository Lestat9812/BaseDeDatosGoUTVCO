package domains

import (
	"time"

	"gorm.io/gorm"
)

type FechasReporte struct {
	gorm.Model
	FechaInicio time.Time  `gorm:"not null"`
	FechaFin    *time.Time `gorm:"default:null"`
}
