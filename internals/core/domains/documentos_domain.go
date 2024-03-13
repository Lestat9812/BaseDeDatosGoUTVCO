package domains

import (
	"time"

	"gorm.io/gorm"
)

type Documento struct {
	gorm.Model
	TipoDocumentoID uint      `gorm:"not null"`
	FechaExpedicion time.Time `gorm:"not null"`
	PersonalID      uint      `gorm:"uniqueIndex;not null"`
	AlumnoID        uint      `gorm:"not null"`
	Pago            bool      `gorm:"not null"`
	Notificacion    Notificacion
	Rechazado       Rechazado
	RegVistos       []Personal `gorm:"many2many:reg_vistos;"` //Relación many to many
	//RegVistos       RegVistos

}
