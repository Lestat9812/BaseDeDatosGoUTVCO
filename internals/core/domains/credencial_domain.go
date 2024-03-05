package domains

import "gorm.io/gorm"

type Credencial struct {
	gorm.Model
	AlumnoID       uint    `gorm:"not null"`
	Nombre         *string `gorm:"default:null"`
	Carrera        *string `gorm:"default:null"`
	Domicilio      *string `gorm:"default:null"`
	CP             *string `gorm:"default:null"`
	CURP           *string `gorm:"default:null"`
	Sangre         *string `gorm:"default:null"`
	Celular        *string `gorm:"default:null"`
	Telefono       *string `gorm:"default:null"`
	Emergencia     *string `gorm:"default:null"`
	TeleEmergencia *string `gorm:"default:null"`
}
