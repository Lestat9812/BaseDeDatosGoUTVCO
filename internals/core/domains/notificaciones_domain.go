package domains

import (
	"time"

	"gorm.io/gorm"
)

func (Notificacion) TableName() string {
	return "notificaciones"
}

type Notificacion struct {
	gorm.Model
	Destino     uint      `gorm:"not null"`
	Autor       uint      `gorm:"not null"`
	DocumentoID uint      `gorm:"not null"`
	Tipo        string    `gorm:"not null"`
	Fecha       time.Time `gorm:"not null"`
	Estado      bool      `gorm:"not null"`
}
