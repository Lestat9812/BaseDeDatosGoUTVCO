package domains

import (
	"gorm.io/gorm"
)

func (EstadoRegistro) TableName() string {
	return "estado_registro"
}

type EstadoRegistro struct {
	gorm.Model
	EstadoD               string `gorm:"not null"`
	Registro              Registro
	ReinscripcionTemporal ReinscripcionTemporal
}
